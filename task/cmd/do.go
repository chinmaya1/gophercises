//package cmd have functionality to implement command line using cobra
package cmd

import (
	"fmt"
	"strconv"

	db "github.com/chinmaya1/gophercises/task/db"
	"github.com/spf13/cobra"
)

//doCmd defines that tasks are completed as well as deleted
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to Parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.RetriveTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTasks(task.Key)

			fmt.Println(id)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
