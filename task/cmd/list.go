//package cmd have functionality to implement command line using cobra
package cmd

import (
	"fmt"

	db "github.com/chinmaya1/gophercises/task/db"

	"github.com/spf13/cobra"
)

//listCmd used to list all the tasks
//Run() implements such a way that inside lists all the tasks stored in Bolt DB
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.RetriveTasks()
		if err != nil {
			fmt.Println("Error displaying tasks:", err.Error())
			return
		}
		if len(tasks) == 0 {
			fmt.Printf("No tasks to finish\n")
			return
		}
		fmt.Println("List of Tasks:")
		for i := 0; i < len(tasks); i++ {
			fmt.Printf("%d.%s\n", i+1, tasks[i].Value)
		}
	},
}

//init() defines add the ListCmd to RootCmd
func init() {
	RootCmd.AddCommand(listCmd)

}
