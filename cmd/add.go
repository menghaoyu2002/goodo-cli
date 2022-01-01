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
	newTodoCmd.SetUsageTemplate(
`Usage: 
  goodo add [new task name]
`)
}

var newTodoCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new task to the TODO list",
	Long: "Add a new task to the TODO list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		numberOfLines := numberOfLines(TODO_FILEPATH)
		file, _ := os.OpenFile(TODO_FILEPATH,  os.O_APPEND|os.O_WRONLY, 0600)
		_, err := file.WriteString(fmt.Sprintf("%d. %s\n", numberOfLines, args[0]));

		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	},
}

// numberOfLines returns the number of lines in the TODO.txt file
func numberOfLines(fileName string) int {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.Count(fileContent, []byte("\n"))
}

