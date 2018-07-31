package cmd

import (
	"fmt"
	"os"

	db "github.com/chinmaya1/gophercises/task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.RetriveTasks()
		if err != nil {
			fmt.Println("Error displaying tasks:", err.Error())
			os.Exit(1)
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

func init() {
	RootCmd.AddCommand(listCmd)

}
