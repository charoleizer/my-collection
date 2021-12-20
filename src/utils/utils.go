package utils

import (
	"regexp"
)

func RemoveSpecialCharacters(text string) string {
	if text == "" {
		return "Empty text"
	}

	//Expression here is fixed. There is no reason to check err
	re, _ := regexp.Compile(`[^\w\s]`)
	text = re.ReplaceAllString(text, "")
	return text
}
