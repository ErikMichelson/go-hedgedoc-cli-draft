package argparse

import (
	"github.com/ErikMichelson/go-hedgedoc/internal/messages"
)

func HandleMainArgs(args []string) {
	switch args[0] {
	case "help":
		HandleHelp(args[1:])
	case "get":

	default:
		messages.PrintUnknownCommand()
	}
}
