/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"os"
	"log"
	"strings"
	"strconv"

	"github.com/spf13/cobra"
)

// swapCmd represents the swap command
var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "swap the order of two tasks in the TODO list",
	Long: `swap the order of two tasks in the TODO list

Example Usage:
  $ goodo
  TODO
  1. First
  2. Second

  $ goodo swap 1 2
  $ goodo
  TODO
  1. Second
  2. First`,
	Args: func (cmd *cobra.Command, args []string) error {
		if (len(args) != 2) {
			return errors.New("requires exactly two line numbers as arguments")
		}
		
		args0Err := isValidLineNumber(args[0])
		args1Err := isValidLineNumber(args[1])
		if args0Err != nil || args1Err != nil {
			return errors.New("both arguments must be valid line numbers")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fileContent, err := os.ReadFile(TODO_FILEPATH)
		if err != nil {
			log.Fatal(err)
		}
		firstIndex, _ := strconv.Atoi(args[0])
		secondIndex, _ := strconv.Atoi(args[1])

		fileLines := strings.Split(string(fileContent), "\n")
		fileLines[firstIndex], fileLines[secondIndex] = increaseLineNumber(fileLines[secondIndex], firstIndex - secondIndex), increaseLineNumber(fileLines[firstIndex], secondIndex - firstIndex)
		swappedFile := strings.Join(fileLines, "\n")

		err = os.WriteFile(TODO_FILEPATH, []byte(swappedFile), 0666)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(swapCmd)
	swapCmd.SetUsageTemplate("Usage:\n  goodo swap [first task number] [second task number]")
}
