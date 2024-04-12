package data

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jsphbtst/spongebobcase/pkg/types"
)

func GetJsonData(fileLoc string) (*types.JsonFile, error) {
	jsonFile, err := os.Open(fileLoc)
	if err != nil {
		return nil, err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var jsonStruct types.JsonFile
	err = json.Unmarshal(jsonData, &jsonStruct)
	if err != nil {
		return nil, err
	}

	return &jsonStruct, nil
}
