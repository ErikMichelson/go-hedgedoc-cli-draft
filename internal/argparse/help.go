package argparse

import (
	"os"

	"github.com/ErikMichelson/go-hedgedoc/internal/messages"
)

func HandleHelp(parameters []string) {
	if len(parameters) < 1 {
		messages.PrintHelp()
		os.Exit(0)
	}
	switch parameters[0] {
	case "post":
		messages.PrintHelpPost()
	case "get":
		messages.PrintHelpGet()
	case "profiles":
		messages.PrintHelpProfiles()
	case "history":
		messages.PrintHelpHistory()
	case "help":
		messages.PrintHelpHelp()
	default:
		messages.PrintHelpUnknown()
	}
	os.Exit(0)
}
