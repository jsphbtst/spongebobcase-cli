package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jsphbtst/spongebobcase/pkg/checkers"
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

/*
 * Would be good to store up to 50 jokes offline in a "jokes.json"
 * file. Pre-fill these with jokes and remove them when new jokes come
 * in later.
 */
func generateFart(cmd *cobra.Command, args []string) {
	offlineText := "hahahaha lmfao farts no internet"
	offlineParsedText := utils.Spongebobify(offlineText)

	hasInternet := checkers.HasInternet()
	if !hasInternet {
		fmt.Println(offlineParsedText)
		os.Exit(0)
	}

	dadJokeStruct, err := dadjokes.Get()
	if err != nil {
		fmt.Println(offlineParsedText)
		os.Exit(0)
	}

	text := dadJokeStruct.Joke
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
