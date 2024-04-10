package utils

import "strings"

func Spongebobify(text string) string {
	alphaMap := GenerateAlphaMap()
	newText := ""
	for i := range text {
		if alphaMap[string(text[i])] {
			newText += strings.ToUpper(string(text[i]))
		} else {
			newText += string(text[i])
		}
	}

	return newText
}
