package lettersorter

import (
	"regexp"
	"sort"
	"strings"
)

//SortString is used to sort a string by vocals and consonants
//ex. SortString("osama") will return "aaoms"
//for sorting more than 20 char its recommended to use SortStringConcurrent()
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

//SortStringConcurrent is used to sort a string by vocals and consonants, implementing concurrency
//ex. SortStringConcurrent("osama") will return "aaoms"
func SortStringConcurrent(text string) string {
	text = strings.ToLower(text)
	vocalsChan, consonantsChan := make(chan []string), make(chan []string)
	go func() {
		vocals := findAll(text, `[aiueo]`)
		sort.Strings(vocals)
		vocalsChan <- vocals
	}()
	go func() {
		consonants := findAll(text, `[bcdfghjklmnpqrstvwxyz]`)
		sort.Strings(consonants)
		consonantsChan <- consonants
	}()
	sorted := append(<-vocalsChan, (<-consonantsChan)...)
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
