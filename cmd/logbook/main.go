package main

import (
	"github.com/N3moAhead/logbook/internal/changelog"
	"github.com/N3moAhead/logbook/internal/git"
)

func main() {
	commits := git.GetCommits()
	changelog.WriteChangelog(commits)
}
