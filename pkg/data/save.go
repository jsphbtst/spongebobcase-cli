package data

import (
	"encoding/json"
	"os"
)

func Save(joke string, fileLoc string) error {
	currentJokesJson, err := GetJsonData(fileLoc)
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
