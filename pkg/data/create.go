package data

import "os"

func CreateJokesJson() error {
	file, err := os.Create("jokes.json")
	if err != nil {
		return err
	}

	defer file.Close()

	content := "{ \"data\": [] }"
	file.Write([]byte(content))
	return nil
}
