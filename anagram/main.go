//An anagram is a word or phrase formed by rearranging the letters of another, using all the original letters exactly once.
// Here's a simple Go program to check if two strings are anagrams

package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	// Remove spaces and convert to lowercase
	s1 = strings.ReplaceAll(strings.ToLower(s1), " ", "")
	s2 = strings.ReplaceAll(strings.ToLower(s2), " ", "")

	// If the lengths of the strings are different, they can't be anagrams
	if len(s1) != len(s2) {
		return false
	}

	// Sort the characters in both strings
	s1Sorted := sortString(s1)
	s2Sorted := sortString(s2)

	// Compare the sorted strings
	return s1Sorted == s2Sorted
}

func sortString(s string) string {
	sRunes := []rune(s)
	sort.Sort(sortRunes(sRunes))
	return string(sRunes)
}

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s1 := "listen"
	s2 := "silent"

	if isAnagram(s1, s2) {
		fmt.Printf("%s and %s are anagrams.\n", s1, s2)
	} else {
		fmt.Printf("%s and %s are not anagrams.\n", s1, s2)
	}
}
