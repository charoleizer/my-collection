package utils

import (
	"regexp"
)

func RemoveSpecialCharacters(text string) string {
	if text == "" {
		return "Empty text"
	}

	re := regexp.MustCompile(`[^\w\s]`)
	text = re.ReplaceAllString(text, "")
	return text
}
