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
	options  = []string{"abc", "df", "option-3"}
	selected = 0 // start with first item as default
	hasQuit  = true
)

func render() {
	// clear the existing screen before rendering new frame
	fmt.Print("\033[2J\033[H") // Clear + home
	for i, item := range options {
		if i == selected {
			utils.TurnText("> "+item, "magentaBright", true, true)
		} else {
			utils.TurnText(item, "magenta", false, false)
		}
	}
}

func List() {
	// store init state
	old_state, err := term.MakeRaw(0)
	if err != nil {
		utils.LogError(err)
		return
	}
	// if pressed quit or ctrl+c for example
	defer term.Restore(0, old_state)

	// listen to events
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// poll
	ticker := time.NewTicker(time.Millisecond * 50)
	defer ticker.Stop()

	// initial render
	//render()

	for !hasQuit {
		select {
		case <-ticker.C:
			key_pressed_msg := make([]byte, 1)
			_, err := os.Stdin.Read(key_pressed_msg)
			if err == nil {

				switch key_pressed_msg[0] {
				case 'q', 3:
					// ctrl+c or quit
					hasQuit = true
				case 10, 13:
					// case enter
					// todo : decide what to do later
				case 65, 75:
					// up arrow
					if selected > 0 {
						selected--
					}

				case 66, 76:
					if selected < len(options)-1 {
						selected++
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
	fmt.Print("\033[?25h") // Show cursor

}
