/*
Copyright © 2021 Menghao Yu menghaoyu2002@gmail.com

*/
package main

import (
	"log"
	"os"

	"github.com/menghaoyu2002/goodo-cli/cmd"
)

func main() {
	initializeTODOFile()
	cmd.Execute()
}

// initializeTODOFile initializes the TODO.txt file in the current directory if it doesn't not exist already
func initializeTODOFile() {
	_, err := os.Open("TODO.txt")
	if err != nil {
		err = os.WriteFile(cmd.TODO_FILEPATH, []byte("TODO"), 0666)
		if err != nil {
			log.Fatal(err)
		}
	} 
}
