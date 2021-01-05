package main

import (
	"os"

	"github.com/ErikMichelson/go-hedgedoc/internal/argparse"
	"github.com/ErikMichelson/go-hedgedoc/internal/config"
	"github.com/ErikMichelson/go-hedgedoc/internal/messages"
)

func main() {
	messages.PrintHeading()

	if !config.Exists() {
		config.DoInitialSetup()
		os.Exit(0)
	}

	config.LoadConfig()
	profile := config.GetCurrentProfile()
	messages.PrintHeadingProfile(profile.Server)

	if len(os.Args) < 2 {
		messages.PrintHelp()
		os.Exit(0)
	}

	argparse.HandleMainArgs(os.Args[1:])
}
