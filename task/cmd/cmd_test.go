package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	db "github.com/chinmaya1/gophercises/task/db"
	"github.com/stretchr/testify/assert"

	homedir "github.com/mitchellh/go-homedir"
)

func setDB() string {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "my.db")
	return dbPath
}

func TestAddCommand(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"Complete gophercises 7"}
	addCmd.Run(addCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Added")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()
	db.Db.Close()
}
func TestRetriveCommand(t *testing.T) {

	dbPath := setDB()
	db.Init(dbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file

	a := []string{"1"}
	listCmd.Run(listCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "List of Tasks:")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	db.Db.Close()
}
func TestDoCommand(t *testing.T) {
	DbPath := setDB()
	db.Init(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"1"}
	doCmd.Run(doCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Marked")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	db.Db.Close()

}

func TestRoot(t *testing.T) {
	err := RootCmd.Execute()
	if err != nil {
		t.Error("error in root command")
	}
}
