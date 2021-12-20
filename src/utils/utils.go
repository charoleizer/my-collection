package utils

import (
	"log"
	"regexp"
)

func RemoveSpecialCharacters(text string) string {
	if text == "" {
		return "Empty text"
	}

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
		return "Invalid text"
	}
	text = re.ReplaceAllString(text, " ")
	return text

}
