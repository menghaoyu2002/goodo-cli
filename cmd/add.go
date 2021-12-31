package cmd

import (
	"bytes"
	"fmt"
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
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// get the number of lines in the file
		fileContent, _ := os.ReadFile(TODO_FILEPATH)
		numberOfLines := bytes.Count(fileContent, []byte("\n"))

		file, _ := os.OpenFile(TODO_FILEPATH,  os.O_APPEND|os.O_WRONLY, 0600)
		_, err := file.WriteString(fmt.Sprintf("%d. %s\n", numberOfLines, args[0]));

		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	},
}

