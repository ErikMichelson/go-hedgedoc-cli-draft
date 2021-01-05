package argparse

import (
	"os"

	"github.com/ErikMichelson/go-hedgedoc/internal/messages"
)

func HandleGet(parameters []string) {
	if len(parameters) < 1 {
		messages.PrintGetNoId()
		os.Exit(0)
	}
	// TODO Add logic
	os.Exit(0)
}
