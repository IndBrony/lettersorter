package lettersorter

import (
	"regexp"
	"sort"
	"strings"
)

//SortString is used to sort a string by vocals and consonants
//ex. SortString("osama") will return "aaoms"
func SortString(text string) string {
	text = strings.ToLower(text)
	vocals, consonants := findAll(text, `[aiueo]`), findAll(text, `[bcdfghjklmnpqrstvwxyz]`)
	sort.Strings(vocals)
	sort.Strings(consonants)
	sorted := append(vocals, consonants...)
	sortedString := ""
	for _, letter := range sorted {
		sortedString += letter
	}
	return sortedString
}

func findAll(text string, regex string) []string {
	compiledRegex := regexp.MustCompile(regex)
	found := compiledRegex.FindAll([]byte(text), -1)
	strings := []string{}
	for _, letter := range found {
		strings = append(strings, string(letter))
	}
	return strings
}
