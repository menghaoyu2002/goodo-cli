/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the n-th tast on the TODO list",
	Long: `Remove the n-th task on the TODO list.

Example Usage:
  $ goodo
  TODO
  1. the first task

  $ goodo remove 1

  $ goodo
  TODO`,
  	Args: func(cmd *cobra.Command, args []string) error {
		// check that exactly one argument is provided
		if len(args) != 1 {
			return errors.New("requires exactly one argument")
		}

		// check that argument is a valud line number
		value, err := strconv.Atoi(args[0])
		if err != nil || value < 1 || value > numberOfLines() {
			return errors.New("argument must be a valid task number")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fileContent, err := os.ReadFile(TODO_FILEPATH)
		if err != nil {
			log.Fatal(err)
		}

		fileLines := strings.Split(string(fileContent), "\n")
		newContent := ""
		for i := 0; i < len(fileLines) - 1; i++ {
			if lineNum, _ := strconv.Atoi(args[0]); i < lineNum {
				newContent += fileLines[i] + "\n"
			} else if (i > lineNum) {
				newContent += IncreaseLineNumber(fileLines[i], -1) + "\n"
			}
		}

		err = os.WriteFile(TODO_FILEPATH, []byte(newContent), 0666)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.SetUsageTemplate(`Usage:
  goodo remove [task number]
`)
}

// ChangeLineNumber increases the line number by lineNumberOffset
func IncreaseLineNumber(line string, lineNumberOffset int) string {
	numbers := strings.Split(line, ".")

	increasedNumber, _ := strconv.Atoi(numbers[0])
	increasedNumber += lineNumberOffset 

	return fmt.Sprint(increasedNumber) + line[len(numbers[0]):]
}
