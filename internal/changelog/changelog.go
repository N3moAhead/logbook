package changelog

import (
	"fmt"
	"os"
	"strings"
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

func WriteChangelog(commits []git.Commit, config Config) {
	var changelogTemplate string = getTemplate(config)
	tmpl, err := template.New("Changelog").Parse(changelogTemplate)
	if err != nil {
		fmt.Println("Error while trying to parse the template: ", err)
		os.Exit(1)
	}

	changelogFile, err := os.Create(config.OutputPath)
	if err != nil {
		fmt.Println("Error while trying to creat an output file: ", err)
		os.Exit(1)
	}
	defer changelogFile.Close()

	err = tmpl.Execute(changelogFile, commits)
	if err != nil {
		panic(err)
	}
}

func getTemplate(config Config) string {
	var templatePath string = config.TemplatePath
	templatePath = strings.TrimSpace(templatePath)
	if templatePath != "" {
		content, err := os.ReadFile(templatePath)
		if err != nil {
			fmt.Println("Error while trying to read template file: ", err)
			return DefaultTemplate
		}
		return string(content)
	}
	return DefaultTemplate
}
