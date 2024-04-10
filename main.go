/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/jsphbtst/spongebobcase/cmd"
	"github.com/jsphbtst/spongebobcase/pkg/data"
)

func main() {
	if _, err := os.Stat("jokes.json"); err != nil {
		err := data.CreateJokesJson()
		if err != nil {
			fmt.Printf("Failed to create jokes file: %s\n", err.Error())
			os.Exit(1)
		}
	}

	jokesJson, err := data.GetJokesJson()
	if err != nil {
		fmt.Printf("Failed to retrieve jokes json file: %s\n", err.Error())
		os.Exit(1)
	}

	cmd.Init(jokesJson.Data)
	cmd.Execute()
}
