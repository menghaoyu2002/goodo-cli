/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const TODO_FILEPATH = "./TODO.txt"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goodo",
	Short: "Goodo is a simple TODO list CLI app",
	Long: "A sweet and simple TODO list CLI app built with Go.\nMy first steps in writing code in this language ❤️",
	Run: func(cmd *cobra.Command, args []string) {
		todoList, _ := os.ReadFile(TODO_FILEPATH)
		fmt.Printf("%s", todoList)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


