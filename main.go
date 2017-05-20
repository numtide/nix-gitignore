package main

import (
	"fmt"
	"log"
	"os"

	"github.com/monochromegane/go-gitignore"
)

const usage = `
Usage: nix-gitignore <gitignore-path> <path> <isDir>

nix-gitignore is a utility that returns the string "true" if the <path> of
type <isDir> is not ignored by the .gitignore files at <gitignore-path>.
Or "false" otherwise.

<isDir> must be "regular" or "directory"

It's only supposed to be used by the accompanying nix library.
`

func main() {
	args := os.Args
	if len(args) != 4 {
		log.Fatalln("invalid number of arguments, should be 3 but got", len(args)-1)
	}

	fmt.Sprintln(os.Stderr, args)

	gitignorePath := args[1]
	path := args[2]
	isDir := false
	switch args[3] {
	case "regular", "symlink", "unknown":
		isDir = false
	case "directory":
		isDir = true
	default:
		log.Fatalln("invalid isDir argument:", args[3])
	}
	if args[3] == "true" {

	}

	matcher, err := gitignore.NewGitIgnore(gitignorePath)
	if err != nil {
		log.Fatalln("gitignore error:", err)
	}

	isMatch := matcher.Match(path, isDir)
	if isMatch {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}
