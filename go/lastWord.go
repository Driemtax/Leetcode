package main

import "strings"

func lengthOfLastWord(s string) int {
	var length int = 0
	s = strings.TrimSpace(s)
	var lastWord string
	lastSpace := strings.LastIndex(s, " ")

	if lastSpace == -1 {
		lastWord = s
	} else {
		lastWord = s[lastSpace+1:]
	}

	length = len(lastWord)

	return length
}
