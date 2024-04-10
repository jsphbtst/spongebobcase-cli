package data

import (
	"encoding/json"
	"os"
)

// TODO: do not store the same text in jokes.json
func SaveJoke(joke string) error {
	currentJokesJson, err := GetJokesJson()
	if err != nil {
		return err
	}

	// TODO: cap at maybe 50 jokes? After which pop the first one,
	// i.e., FIFO
	currentJokesJson.Data = append(currentJokesJson.Data, joke)
	configsBytes, err := json.MarshalIndent(currentJokesJson, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("jokes.json", configsBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
