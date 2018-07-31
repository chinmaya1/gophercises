package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chinmaya1/gophercises/task/cmd"
	"github.com/chinmaya1/gophercises/task/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	check(db.Init(dbPath))
	check(cmd.RootCmd.Execute())
}
func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
