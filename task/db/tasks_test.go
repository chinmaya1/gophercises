package db

import (
	"fmt"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestMain(t *testing.M) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	Init(dbPath)
	t.Run()
}

func TestAddTask(t *testing.T) {
	_, err := AddTask("Task is added sucessfully")
	if err != nil {
		t.Errorf("Task is not added: %d", err)

	}

}

func TestRetriveTask(t *testing.T) {
	var tasks []Task
	tasks, err := RetriveTasks()
	if len(tasks) == 0 {
		t.Errorf("Error in retiriving :%d", err)
	}
}

func TestDeleteTask(t *testing.T) {
	err := DeleteTasks(1)
	fmt.Println(err)
	if err != nil {
		t.Error("Error in Deleteing tasks")
	}
}

func TestIntToByte(t *testing.T) {
	bytes := inttobyte(10)
	fmt.Println(bytes)
	if bytes == nil {
		t.Errorf("Error for converting int to byte : %d ", bytes)

	}
}

func TestByteToInt(t *testing.T) {
	var s []byte
	s = make([]byte, 8, 12)
	s = []byte{0, 0, 0, 0, 0, 0, 0, 10}
	val := bytetoint(s)
	if val == 0 {
		t.Error("Expected int got", val)
	}
}
