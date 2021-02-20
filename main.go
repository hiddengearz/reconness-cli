package main

import (
	"os"

	"github.com/hiddengearz/reconness-cli/cmd"
	"github.com/hiddengearz/reconness-cli/internal"
)

func main() {
	stat, _ := os.Stdin.Stat() //Cheks if any data is being piped in from other commands
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		internal.Pipe = true
	} else {
		internal.Pipe = false
	}

	cmd.Execute()
}
