//package cmd have functionality to implement command line using cobra
package cmd

import (
	"fmt"
	"strings"

	db "github.com/chinmaya1/gophercises/task/db"
	"github.com/spf13/cobra"
)

//addCmd add tasks to the list
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",

	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		msg := "task is not added"
		_, err := db.AddTask(task)
		if err == nil {
			msg = fmt.Sprintf("Added \"%s\"\n", task)
		}
		fmt.Print(msg)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
