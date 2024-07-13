package main

import (
	"fmt"

	"github.com/N3moAhead/logbook/internal/changelog"
	"github.com/N3moAhead/logbook/internal/git"
)

func main() {
	fmt.Println("The beginning of logbook.")
	commits := git.GetCommits()
	for _, value := range commits {
		fmt.Printf("Author: %s\nSubject: %s\n\n", value.Author, value.Subject)
	}
	changelog.WriteChangelog()
}
