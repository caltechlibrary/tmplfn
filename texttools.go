package tmplfn

import (
	"strings"
	"text/template"
)

var (
	TextTools = template.FuncMap{
		"slug":          slug,
		"unslug":        unslug,
		"english_title": englishTitle,
	}
)

// slug - takes a string and converts it be filename friendly
func slug(s string) string {
	return strings.ToLower(strings.Replace(strings.Replace(strings.Replace(s, "-", "_", -1), " ", "-", -1), "/", "~", -1))
}

// unslug takes a filename and converts into a title friendly string
func unslug(s string) string {
	return strings.Replace(strings.Replace(strings.Replace(s, "-", " ", -1), "_", "-", -1), "~", "/", -1)
}

// englishTitle - uses an improve capitalization rules for English titles.
// This is based on the approach suggested in the Go language Cookbook:
//     http://golangcookbook.com/chapters/strings/title/
func englishTitle(s string) string {
	words := strings.Fields(s)
	smallwords := " a an on the to of in "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && index != 0 {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
