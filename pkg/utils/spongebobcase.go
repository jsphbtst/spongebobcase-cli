package utils

import "strings"

func Spongebobify(text string) string {
	alphaMap := GenerateAlphaMap()

	var newText strings.Builder
	for _, runeChar := range text {
		char := string(runeChar)
		if alphaMap[char] {
			newText.WriteString(strings.ToUpper(char))
			continue
		}

		newText.WriteString(char)
	}

	return newText.String()
}
