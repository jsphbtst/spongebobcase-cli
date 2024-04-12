package data

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jsphbtst/spongebobcase/pkg/types"
)

func GetJokesJson(fileLoc string) (*types.JsonFile, error) {
	jokesJsonFile, err := os.Open(fileLoc)
	if err != nil {
		return nil, err
	}

	jokesJsonData, err := io.ReadAll(jokesJsonFile)
	if err != nil {
		return nil, err
	}

	var jokesJson types.JsonFile
	err = json.Unmarshal(jokesJsonData, &jokesJson)
	if err != nil {
		return nil, err
	}

	return &jokesJson, nil
}
