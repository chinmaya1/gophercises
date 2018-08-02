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

	record, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"Complete gophercises 7"}
	addCmd.Run(addCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Added")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Db.Close()
}
func TestRetriveCommand(t *testing.T) {

	dbPath := setDB()
	db.Init(dbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{""}
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

func TestListCommandNegative(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "testing.db")
	db.Init(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{""}
	listCmd.Run(listCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "No tasks to finish")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	db.Db.Close()

}

func TestListCmdInvoke(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "my.db")
	db.Init(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{""}
	db.Db.Close()
	listCmd.Run(listCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Error displaying tasks:")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()

}
func TestDoNegDbCommand(t *testing.T) {
	DbPath := setDB()
	db.Init(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"1"}
	db.Db.Close()
	doCmd.Run(doCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Something went wrong:")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()

}

func TestDoCmdNegative(t *testing.T) {
	DbPath := setDB()
	db.Init(DbPath)

	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"100000000"}
	doCmd.Run(doCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Invalid task number")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	db.Db.Close()

}

func TestDoCmdInvalid(t *testing.T) {
	DbPath := setDB()
	db.Init(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"b"}
	doCmd.Run(doCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Failed to Parse the argument:")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	db.Db.Close()

}
