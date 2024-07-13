package changelog

import (
	"fmt"
	"os"

	"github.com/N3moAhead/logbook/internal/git"
)

func WriteChangelog(commits []git.Commit) {
	f, err := os.Create("CHANGELOG.md")
	if err != nil {
		fmt.Println("Error while trying to create a file: ", err)
		os.Exit(1)
	}
	defer f.Close()

	f.WriteString("# Changelog\n\n")
	for _, commit := range commits {
		_, err := f.WriteString("- " + commit.Subject + "\n")
		if err != nil {
			fmt.Println("Error while trying to write to the ouput file: ", err)
			os.Exit(1)
		}
	}
}
