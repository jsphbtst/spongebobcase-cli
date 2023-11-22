package cmd

import (
	"fmt"
	"strings"

	"github.com/jsphbtst/spongebobcase/pkg/utils"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate spongebob case text",
	Long: `Generate spongebob case text in one command.
	For example:

	spongebobcase generate -t "this is a sample text"
`,
	Run: generateSpongebobCase,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("text", "t", "just pick one", "The text that needs to be spongebobcasified")
}

func generateSpongebobCase(cmd *cobra.Command, args []string) {
	text, _ := cmd.Flags().GetString("text")

	alphaMap := utils.GenerateAlphaMap()
	newText := ""
	for i := range text {
		if alphaMap[string(text[i])] {
			newText += strings.ToUpper(string(text[i]))
		} else {
			newText += string(text[i])
		}
	}

	fmt.Println(newText)
}
