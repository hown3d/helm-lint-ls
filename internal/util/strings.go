package util

import "strings"

// BetweenStrings gets the substring between two strings.
func BetweenStrings(value string, leftDelimiter string, rightDelimiter string) (v string, present bool) {
	posFirst := strings.Index(value, leftDelimiter)
	if posFirst == -1 {
		return value, false
	}
	posLast := strings.Index(value, rightDelimiter)
	if posLast == -1 {
		return value, false
	}
	posFirstAdjusted := posFirst + len(leftDelimiter)
	if posFirstAdjusted >= posLast {
		return value, false
	}
	return value[posFirstAdjusted:posLast], true
}

// AfterStrings gets the substring after a string.
func AfterStrings(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}
