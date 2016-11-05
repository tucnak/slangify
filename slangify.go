package slangify

import (
	"regexp"
	"strings"
)

// Slangify ...
func Slangify(text string) string {
	text = strings.ToLower(text)

	for what, with := range British {
		text = strings.Replace(text, what, with, -1)
	}

	for what, with := range Shortcuts {
		r := regexp.MustCompile("\\b" + what + "\\b")
		text = r.ReplaceAllLiteralString(text, with)
	}

	return text
}
