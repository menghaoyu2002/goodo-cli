package cmd

import (
	//"fmt"
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newTodoCmd)
}

var newTodoCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new task to the TODO list",
	Args: func (cmd *cobra.Command, args []string) error {
		if (len(args) > 1) {
			return errors.New("too many arguments")
		} else if (len(args) == 0) {
			return errors.New("task name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := os.OpenFile("./TODO.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if _, err := file.WriteString("\n" + args[0]); err != nil {
			log.Fatal(err)
		}
		file.Close()
	},
}

