/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/jsphbtst/spongebobcase/cmd"
)

type JokesJsonFile struct {
	Data []string `json:"data"`
}

func main() {
	// Jokes file
	if _, err := os.Stat("jokes.json"); err != nil {
		file, err := os.Create("jokes.json")
		if err != nil {
			log.Fatalf("Failed to created jokes file: %s\n", err.Error())
			os.Exit(1)
		}

		content := "{ \"data\": [] }"
		file.Write([]byte(content))
		defer file.Close()
	}

	jokesJsonFile, err := os.Open("jokes.json")
	if err != nil {
		log.Fatalf("Failed to open ∆okes file: %s\n", err.Error())
		os.Exit(1)
	}

	jokesJsonData, err := io.ReadAll(jokesJsonFile)
	if err != nil {
		log.Fatalf("Failed to parse configs file: %s\n", err.Error())
		os.Exit(1)
	}

	var jokesJson JokesJsonFile
	err = json.Unmarshal(jokesJsonData, &jokesJson)
	if err != nil {
		log.Fatalf("Failed to unmarshal configs file: %s\n", err.Error())
		os.Exit(1)
	}

	cmd.Init(jokesJson.Data)
	cmd.Execute()
}
