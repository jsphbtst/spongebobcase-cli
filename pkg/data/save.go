package data

import (
	"encoding/json"
	"os"
)

// TODO: do not store the same text in cache.json
func Save(joke string, fileLoc string) error {
	currentJokesJson, err := GetJokesJson(fileLoc)
	if err != nil {
		return err
	}

	currentJokesJson.AddData(joke)
	configsBytes, err := json.MarshalIndent(currentJokesJson, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileLoc, configsBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
