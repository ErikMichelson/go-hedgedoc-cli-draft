package messages

import "fmt"

func PrintHeading() {
	fmt.Println("HedgeDoc CLI version 2.0")
}

func PrintUnknownCommand() {
	fmt.Println("Unknown command entered.")
}

func PrintFirstUse() {
	fmt.Println(`
Welcome! It looks as this is the first time you're using
the hedgedoc-cli. Let's get started with some quick setup.

If you need help with this tool, call "hedgedoc-cli help".`)
}

func PrintHeadingProfile(server string) {
	fmt.Println("Server: ", server)
}
