package list

// steps :
// 1. clear screen, render init/updated list
// 2. change from canonical to raw mode in terminal (so u can listen to events like up/down arrows)
// 3. listen to key events (do a poll every X ms)
// 4. save the new state based on key event
// 5. move to step-1

import (
	"cute-line-interface/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/term"
)

var (
	options      = []string{}
	selected     = 0 // start with first item as default
	hasQuit      = false
	prevSelected = -1 // used so that we repaint only last + new selection, not whole list
	startLine    = 0  // FIX for bug : spaces shown above and below list rendering
)

func render() {
	// FIRST TIME ONLY: Full initial render
	if prevSelected == -1 {
		//fmt.Print("\033[2J\033[H") // Clear + home (once)
		// above clear logic has been removed as we start from wherever the cursor was last time
		for i := range options {
			prefix := "  "
			if i == selected {
				prefix = "> "
			}
			color := "magenta"
			if i == selected {
				color = "magentaBright"
			}
			fmt.Println(utils.TurnText(prefix+options[i], color, i == selected, i == selected))
		}
		fmt.Println(utils.TurnText("↑↓ navigate, Enter select, q quit", "cyanBright", false, false))
		prevSelected = selected
		return
	}

	// Save cursor, move up to start of list, repaint changed lines, restore cursor
	totalLines := len(options) + 1 // +1 for help text

	// Move cursor up to start of list
	// \033[1A would move cursor 1 up
	// so after rendering list of 5 items, cursor is now 6 lines below first item
	// so move it up by 6 positions
	fmt.Printf("\033[%dA", totalLines)

	// ONLY repaint changed lines (NO full list redraw)
	// \033[1B moves cursor one down
	// so move down till you reach prev selected item
	for i := 0; i < prevSelected; i++ {
		fmt.Printf("\033[1B")
	}

	// after reaching last selection \033[2K to clear entire line
	// /r to go back to start of the same line
	// then repaint to un-highlighted style
	fmt.Printf("\033[2K\r")
	fmt.Printf("%s\n", utils.TurnText("  "+options[prevSelected], "magenta", false, false))

	// note the kast print gave us a \n so we have already moved one line down
	// Move to new selection line
	// if prev = 3, new = 5, then move 5-3-1 (1 cause we are already one down via \n)
	linesDiff := selected - prevSelected - 1
	// it says if old selection is below, then move cursor up, else down
	if linesDiff > 0 {
		fmt.Printf("\033[%dB", linesDiff) // Move down
	} else if linesDiff < 0 {
		fmt.Printf("\033[%dA", -linesDiff) // Move up
	}
	fmt.Printf("\033[2K\r") // Clear line, return to start
	// paint new selection
	fmt.Printf("%s\n", utils.TurnText("> "+options[selected], "magentaBright", true, true))

	// Move cursor to bottom (after help text)
	remainingLines := len(options) - selected
	fmt.Printf("\033[%dB", remainingLines)

	prevSelected = selected
}

func List(list_items []string) {
	options = list_items
	hasQuit = false
	prevSelected = -1

	// fails here on WINDOWS fix: use fd value instead of '0'
	fd := int(os.Stdin.Fd())

	// store init state
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		utils.LogError(err)
		return
	}
	// if pressed quit or ctrl+c for example, use fd here as well
	defer term.Restore(fd, oldState)

	// listen to events
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)

	// poll
	ticker := time.NewTicker(time.Millisecond * 50)
	defer ticker.Stop()

	// Hide cursor during selection
	fmt.Print("\033[?25l")

	// initial render
	render()

	for !hasQuit {
		select {
		case <-ticker.C:
			// NON-BLOCKING: timeout so doesn't hang
			os.Stdin.SetReadDeadline(time.Now().Add(time.Millisecond))
			b := make([]byte, 3) // Arrow keys = ESC [ A/B (3 bytes)
			n, err := os.Stdin.Read(b)
			if n > 0 && err == nil {
				switch b[0] {
				case 'q', 3: // q or Ctrl+C byte
					// Move cursor to bottom and show it
					fmt.Printf("\033[%dB", len(options)-selected)
					fmt.Print("\033[?25h")
					hasQuit = true
				case 10, 13: // Enter
					// Move cursor to bottom, show it, print selection
					fmt.Printf("\033[%dB", len(options)-selected)
					fmt.Print("\033[?25h")
					fmt.Printf("\nSelected: %q\n", options[selected])
					hasQuit = true
				case 27: // ESC - start of arrow sequence
					if n >= 3 && b[1] == '[' {
						switch b[2] {
						case 'A': // Up arrow
							if selected > 0 {
								selected--
							}
						case 'B': // Down arrow
							if selected < len(options)-1 {
								selected++
							}
						}
					}
				}
				// render @ end of each cycle (only if not quitting)
				if !hasQuit {
					render()
				}
				// // render @ end of each cycle
				// render()
			}
		case <-sig:
			hasQuit = true
		}
	}

	// Just show cursor again, don't clear screen
	fmt.Print("\033[?25h")
}
