package git

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Commit struct {
	Hash          string
	Subject       string
	Body          string
	Author        string
	AuthorEmail   string
	Commiter      string
	CommiterEmail string
}

func GetCommits() []Commit {
	var gitLog string = getGitLog()
	var commitStrings []string = getCommitStrings(gitLog)
	var commits []Commit = getCommitStructs(commitStrings)
	return commits
}

func getGitLog() string {
	cmd := exec.Command(
		"git",
		"log",
		"--pretty='hash':{%H},'subject':{%s},'body':{%b},'author':{%an},'authorEmail':{%ae},'committer':{%cn},'committerEmail':{%ce};",
	)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	return string(output)
}

func getCommitStrings(gitLog string) []string {
	var commits []string = strings.Split(gitLog, ";")
	return commits
}

func getCommitStructs(commitStrings []string) []Commit {
	var commits []Commit
	for _, commitString := range commitStrings {
		commitString = strings.TrimSpace(commitString)
		if commitString == "" {
			continue
		}
		commits = append(commits, parse2CommitStruct(commitString))
	}
	return commits
}

func parse2CommitStruct(commitString string) Commit {
	var commit Commit
	regex := "'.*?':{.*?}"
	re, err := regexp.Compile(regex)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	matches := re.FindAllString(commitString, -1)
	for _, match := range matches {
		key, value := parsePair(match)
		switch key {
		case "hash":
			commit.Hash = value
		case "subject":
			commit.Subject = value
		case "body":
			commit.Body = value
		case "author":
			commit.Author = value
		case "authorEmail":
			commit.AuthorEmail = value
		case "committer":
			commit.Commiter = value
		case "committerEmail":
			commit.CommiterEmail = value
		default:
			fmt.Println("Error: Undefined key: ", key, " in parse2CommitStruct")
			os.Exit(1)
		}
	}
	return commit
}

func parsePair(pair string) (string, string) {
	regex := "'(.*?)':{(.*)}"
	re, err := regexp.Compile(regex)
	if err != nil {
		fmt.Printf("Error parsing regex %s\n", err)
	}
	matches := re.FindStringSubmatch(pair)
	return matches[1], matches[2]
}
