package changelog

import (
	"os"
	"text/template"

	"github.com/N3moAhead/logbook/internal/git"
)

const DefaultTemplate string = `# Changelog

## Commits
{{ range . }}
- 🔧 **{{ .Subject }}** 
  - 📅 **Hash**: {{ .Hash }}
  - 👤 **Author**: {{ .Author }} ({{ .AuthorEmail }})
  - ✍️ **Commiter**: {{ .Commiter }} ({{ .CommiterEmail }})
  - 📝 **Details**: 
    {{ .Body }}{{ end }}
`

func WriteChangelog(commits []git.Commit) {
	tmpl, err := template.New("Changelog").Parse(DefaultTemplate)
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
