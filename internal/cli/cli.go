package cli

import (
	"os"
)

func HandleCommand() {
	if len(os.Args) < 2 {
		println("No command provided")
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "set-role":
		setRole()
	case "user-details":
		printUserDetails()
	case "prelude-content":
		preludeContent()
	default:
		println("Unknown command")
		os.Exit(1)
	}
}
