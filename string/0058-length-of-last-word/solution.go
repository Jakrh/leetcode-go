package solution

import (
	"regexp"
	"strings"
)

var re1 = regexp.MustCompile(`[A-Za-z]+\s*$`)

func lengthOfLastWord1(s string) int {
	match := re1.Find([]byte(s))
	return len(strings.TrimRight(string(match), " "))
}

// Example:
// "   fly me   to   the moon  "
// to "moon  "

// ---------------

var re2 = regexp.MustCompile(`[A-Za-z]+`)

func lengthOfLastWord2(s string) int {
	matches := re2.FindAllString(s, -1)
	return len(matches[len(matches)-1])
}

// Example:
// "   fly me   to   the moon  " to
// ["fly", "me", "to", "the", "moon"]

// ---------------

func lengthOfLastWord3(s string) int {
	end := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && end == 0 {
			end = i
		} else if s[i] == ' ' && end != 0 {
			return end - i
		}
	}

	return end + 1
}
