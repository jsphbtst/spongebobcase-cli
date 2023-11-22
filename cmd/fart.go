package cmd

import (
	"fmt"
	"strings"

	"github.com/jsphbtst/spongebobcase/pkg/dadjokes"
	"github.com/jsphbtst/spongebobcase/pkg/utils"
	"github.com/spf13/cobra"
)

var fartCmd = &cobra.Command{
	Use:   "fart",
	Short: "Fart out spongebob case text",
	Long: `Fart out spongebob case text in one command.
	For example:

	spongebobcase fart
`,
	Run: generateFart,
}

func init() {
	rootCmd.AddCommand(fartCmd)
}

func generateFart(cmd *cobra.Command, args []string) {
	var text string

	dadJokeStruct, err := dadjokes.Get()
	if err != nil {
		text = "hahahaha lmfao farts no internet"
	} else {
		text = dadJokeStruct.Joke
	}

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
