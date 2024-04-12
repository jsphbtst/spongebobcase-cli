package data

import (
	"fmt"
	"os"
)

func CreateJokesJson(path string, filename string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	fileLoc := fmt.Sprintf("%s/%s", path, filename)
	file, err := os.Create(fileLoc)
	if err != nil {
		return err
	}

	defer file.Close()

	content := "{ \"data\": [] }"
	file.Write([]byte(content))
	return nil
}
