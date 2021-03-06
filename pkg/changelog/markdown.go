package changelog

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hekike/unchain/pkg/parser"
)

// Generate generate markdown output
func Generate(version string, commits []parser.ConventionalCommit) string {
	var out bytes.Buffer
	var patch = false
	var minor = false
	var major = false

	// Tag Header
	date := time.Now().Format("2006-01-02")
	out.WriteString(fmt.Sprintf("<a name=\"%s\"></a>\n", version))
	out.WriteString(fmt.Sprintf("## %s (%s)\n\n\n", version, date))

	// Patch
	for _, commit := range commits {
		if commit.SemVerChange == parser.Patch &&
			// Skip non user facing commits from changelog
			commit.Type != "test" && commit.Type != "chore" && commit.Type != "refactor" {

			if patch == false {
				out.WriteString("#### Bug Fixes\n")
			}
			out.WriteString(getCommitLine(&commit))
			patch = true
		}
	}
	if patch == true {
		out.WriteString("\n")
	}

	// Minor
	for _, commit := range commits {
		if commit.SemVerChange == parser.Minor {
			if minor == false {
				out.WriteString("\n#### Features\n")
			}
			out.WriteString(getCommitLine(&commit))
			minor = true
		}
	}
	if minor == true {
		out.WriteString("\n")
	}

	// Major
	for _, commit := range commits {
		if commit.SemVerChange == parser.Major {
			if major == false {
				out.WriteString("\n#### Breaking Changes\n")
			}
			out.WriteString(getBreakingLine(&commit))
			major = true
		}
	}
	if major == true {
		out.WriteString("\n")
	}

	// No user facing commit
	if patch == false && minor == false && major == false {
		out.WriteString("* There is no user facing commit in this version\n")
	}

	out.WriteString("\n\n")

	return out.String()
}

func getCommitLine(commit *parser.ConventionalCommit) string {
	var out bytes.Buffer

	out.WriteString("\n* ")
	if len(commit.Component) > 0 {
		c := fmt.Sprintf("**%s:** ", commit.Component)
		out.WriteString(c)
	}
	out.WriteString(commit.Description)
	out.WriteString(" ")
	out.WriteString(commit.Hash)

	return out.String()
}

func getBreakingLine(commit *parser.ConventionalCommit) string {
	var out bytes.Buffer

	out.WriteString("\n* ")
	out.WriteString(commit.Breaking)
	out.WriteString(" ")
	out.WriteString(commit.Hash)

	return out.String()
}
