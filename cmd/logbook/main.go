package main

import (
	"fmt"

	"github.com/N3moAhead/logbook/internal/changelog"
	"github.com/N3moAhead/logbook/internal/git"
)

func main() {
	fmt.Println("The beginning of logbook.")
	changelog.WriteChangelog()
	git.GetCommits()
}
