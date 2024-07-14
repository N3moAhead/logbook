package changelog

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/N3moAhead/logbook/internal/git"
)

func WriteChangelog(commits []git.Commit) {
	var templatePath = filepath.Join("templates", "changelog.tmpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	changelogFile, err := os.Create("CHANGELOG.md")
	if err != nil {
		panic(err)
	}
	defer changelogFile.Close()

	err = tmpl.Execute(changelogFile, commits)
	if err != nil {
		panic(err)
	}
}
