package config

import (
	"fmt"
	"net/url"

	"github.com/ErikMichelson/go-hedgedoc/internal/messages"
)

func DoInitialSetup() {
	messages.PrintFirstUse()
	DoInterview()
}

func DoInterview() {
	var server, token string
	fmt.Println(`
Please enter the URL of your HedgeDoc instance below.
Example: https://demo.hedgedoc.org`)
	fmt.Print(" > ")
	num, err := fmt.Scanln(&server)
	if err != nil {
		fmt.Println("This seems to not have worked.")
		DoInterview()
		return
	}
	_, err = url.ParseRequestURI(server)
	if err != nil {
		fmt.Println("You probably mistyped. Be sure to include the https:// at the beginning.")
		DoInterview()
		return
	}
	fmt.Println(`
Okay. Now enter an access token for your account below.
You can generate access tokens on your profile page.`)
	fmt.Print(" > ")
	num, err = fmt.Scanln(&token)
	if err != nil || num != 1 {
		fmt.Println("This seems to not have worked.")
		DoInterview()
		return
	}
	SaveProfile(Profile{0, server, token})
	fmt.Println(`
Great! Your credentials are saved.
For later changes, you may edit them via the
"hedgedoc-cli profiles" action.`)
}
