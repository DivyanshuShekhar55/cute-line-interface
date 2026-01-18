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
)

func render() {
    // FIRST TIME ONLY: Full initial render
    if prevSelected == -1 {
        fmt.Print("\033[2J\033[H") // Clear + home (once)
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

    // ONLY repaint changed lines (NO full list redraw)
    // Fix old line (remove highlight)
    fmt.Printf("\033[%d;1H\033[2K", prevSelected+1) // Cursor to old line # (1-based), clear line
    fmt.Printf("%s\n", utils.TurnText("  "+options[prevSelected], "magenta", false, false))

    // Highlight new line  
    fmt.Printf("\033[%d;1H\033[2K", selected+1) // Cursor to new line # (1-based), clear line
    fmt.Printf("%s\n", utils.TurnText("> "+options[selected], "magentaBright", true, true))

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

    // initial render
    render()

    for !hasQuit {
        select {
        case <-ticker.C:
            // NON-BLOCKING: timeout so doesn't hang
            os.Stdin.SetReadDeadline(time.Now().Add(time.Millisecond))
            b := make([]byte, 3)  // Arrow keys = ESC [ A/B (3 bytes)
            n, err := os.Stdin.Read(b)
            if n > 0 && err == nil {
                switch b[0] {
                case 'q', 3:  // q or Ctrl+C byte
                    hasQuit = true
                case 10, 13:  // Enter
                    fmt.Printf("\nSelected: %q\n", options[selected])
                    hasQuit = true
                case 27:  // ESC - start of arrow sequence
                    if n >= 3 && b[1] == '[' {
                        switch b[2] {
                        case 'A':  // Up arrow
                            if selected > 0 {
                                selected--
                            }
                        case 'B':  // Down arrow
                            if selected < len(options)-1 {
                                selected++
                            }
                        }
                    }
                }
                // render @ end of each cycle
                render()
            }
        case <-sig:
            hasQuit = true
        }
    }

    // after quit cursor might not show up so fix here
    fmt.Print("\033[?25h\033[2J") // Show cursor + final clear
}
