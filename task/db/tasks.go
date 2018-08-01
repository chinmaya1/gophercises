//package db consists of functionality to perform add ,retrive and delete operation using Bolt DB
package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

//Varible taskBucket create bucket for tasks
var taskBucket = []byte("tasks")

//Transaction parameter
var Db *bolt.DB

//Task declare as struct
type Task struct {
	Key   int
	Value string
}

//Init method initialises Database and creates task bucket if doesn't exist.
func Init(dbPath string) error {
	var err error
	Db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

//AddTask adds tasks to the Task List
func AddTask(task string) (int, error) {
	var id int
	err := Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := inttobyte(id)
		return b.Put(key, []byte(task))
	})

	return id, err
}

//Retrive retrives data from Bolt DB and returns a list of tasks
func RetriveTasks() ([]Task, error) {
	var tasks []Task
	err := Db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket(taskBucket)
		c := bk.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   bytetoint(k),
				Value: string(v),
			})
		}
		return nil
	})

	return tasks, err
}

//DeleteTasks delete tasks by id
func DeleteTasks(key int) error {
	return Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(inttobyte(key))
	})
}

//inttobyte converts integer to byte
func inttobyte(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//bytetoint converts bytes into integer
func bytetoint(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
