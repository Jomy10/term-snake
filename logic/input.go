package snake

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

// blocking
func ListenDirection(channel chan Direction) {
	// term.RawMode(t)
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(fmt.Sprintf("ERROR: %s", err))
		}

		// Clear channel
		select {
		case <-channel:
		default:
		}

		switch char {
		case 'z', 'Z', 'w', 'W':
			channel <- Up
		case 'a', 'A', 'q', 'Q':
			channel <- Left
		case 's', 'S':
			channel <- Down
		case 'd', 'D':
			channel <- Right
		case 'c', 'C':
			os.Exit(0)
		}
	}
}
