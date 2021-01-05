package messages

import "fmt"

func PrintHelp() {
	fmt.Println(`
Usage: hedgedoc-cli <action> [parameters...]

Available actions:
post      Pushes existing content to the server
get       Fetches note content from the server
history   View a list of your recent notes
profiles  Manage user-accounts connected with this CLI
help      Get help for the usage of actions

Type "hedgedoc-cli help <action>" for more details on the actions.`)
}

func PrintHelpPost() {
	fmt.Println(`
Usage: hedgedoc-cli post [filename] [alias]

The post action pushes local markdown documents to the server.
You can either pipe the markdown-content into this command or
specify the file to import explicitly.

If your server supports "Free-URL"-mode, you may define a custom
alias under which the file should be imported.

Examples:
hedgedoc-cli post /tmp/test.md
  -- The file test.md located in /tmp will be imported to the server.

cat /tmp/test.md | hedgedoc-cli post
  -- Same as above, but this time you're piping your input into the
     cli instead of setting an explicit file name. 

hedgedoc-cli post /tmp/test.md my-test-note
  -- If your server supports "Free-URL"-mode, a new note will be created
     with the alias "my-test-note" with the content of the file test.md.`)
}

func PrintHelpGet() {
	fmt.Println(`
Usage: hedgedoc-cli get <note-id> [output file | -s]

The get action downloads the markdown content of a note.
You can either specify an output file, in which the content will be saved
or the CLI will determine the filename by its own.

If you set the parameter "-s", the note content will be shown in the
terminal instead of saving it into a file.

Examples:
hedgedoc-cli get my-test-note
  -- Fetches the note with the alias "my-test-note" and stores it stores
     it into a file "my-test-note.md".

hedgedoc-cli get 9NVgJD68SdSBx_z-hUO5cQ test.md
  -- Fetches the note with the given id and stores it into "test.md".

hedgedoc-cli get 9NVgJD68SdSBx_z-hUO5cQ -s
  -- Fetches the note with the given id and shows it in the terminal
     instead of storing it to a file.`)
}

func PrintHelpProfiles() {
	fmt.Println(`
Usage: hedgedoc-cli profiles <command> [parameters]

The profiles action allows to add, change and remove user accounts
connected to this CLI. It also allows to switch the active profile.

Following sub-commands are available:
list     List all stored profiles
add      Add a new profile and switch to it
remove   Removes a profile from the list
switch   Switches to another profile

The sub-commands remove and switch take the number of the profile to
operate on as parameter. The profile numbers are displayed with the
list sub-command.`)
}

func PrintHelpHistory() {
	fmt.Println(`
Usage: hedgedoc-cli history <command> [parameters]`)
}

func PrintHelpHelp() {
	fmt.Println(`
Usage: hedgedoc-cli help <action>

The help action displays more information about the usage and
examples for the given action.

Example:
hedgedoc-cli help post
  -- Displays the help for the post-action.`)
}

func PrintHelpUnknown() {
	fmt.Println(`There is no such action.
Try "hedgedoc-cli" help for a list of available actions.`)
}
