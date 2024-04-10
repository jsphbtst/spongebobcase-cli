package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jsphbtst/spongebobcase/pkg/utils"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate spongebob case text",
	Long: `Generate spongebob case text in one command.
	For example:

	spongebobcase generate "this is a sample text"
`,
	Run: generateSpongebobCase,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateSpongebobCase(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Incorrect usage")
		os.Exit(1)
	}

	text := args[0]
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
