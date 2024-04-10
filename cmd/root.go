/*
 * Copyright © 2024 Joseph Bautista jsphbtst@proton.me
 */

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spongebobcase",
	Short: "spongebobcase",
	Long:  "spongebobcase",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
