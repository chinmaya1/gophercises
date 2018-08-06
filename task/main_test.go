package main

import (
	"testing"

	"github.com/CloudBroker/dash_utils/dashtest"
)

func TestM(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("following error occured while the main function was executed : ")
		}
	}()
	main()
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
	m.Run()

}
