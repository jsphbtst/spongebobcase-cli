/*
 * Copyright © 2023 Joseph Bautista jsphbtst@proton.me
 */
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jsphbtst/spongebobcase/cmd"
	"github.com/jsphbtst/spongebobcase/pkg/data"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	rootPwd := filepath.Join(home, ".spongebobcase")
	filename := "data.json"
	fileLoc := fmt.Sprintf("%s/%s", rootPwd, filename)
	if _, err := os.Stat(fileLoc); err != nil {
		err := data.CreateJokesJson(rootPwd, filename)
		if err != nil {
			fmt.Printf("Failed to create jokes file: %s\n", err.Error())
			os.Exit(1)
		}
	}

	jokesJson, err := data.GetJsonData(fileLoc)
	if err != nil {
		fmt.Printf("Failed to retrieve jokes json file: %s\n", err.Error())
		os.Exit(1)
	}

	cmd.Init(jokesJson.Data)
	cmd.InitFileLoc(fileLoc)
	cmd.Execute()
}
