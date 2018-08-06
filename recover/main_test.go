package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testhandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test entered test handler, this should not happen")
	}
	return http.HandlerFunc(fn)
}

func TestRecoveryM(t *testing.T) {
	handler := http.HandlerFunc(PanicHandler)
	executeRequest("Get", "/panic", RecoveryMw(handler))

}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}

func TestSourceCodeHandler(t *testing.T) {

	ts := httptest.NewServer(GetHandler())
	defer ts.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/chinmaya/go/src/github.com/chinmaya1/gophercises/recover/main.go&line=24", nil), status: 200},
		{name: "2: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/chinmaya/go/src/github.com/chinmaya1/gophercises/recoer/main.go&line=24", nil), status: 500},
		{name: "2: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/chinmaya/go/src/github.com/chinmaya1/gophercises/recover/main.go&line=et", nil), status: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}

func TestMakeLinks(t *testing.T) {
	str := `
	goroutine 6 [running]:
runtime/debug.Stack(0xc42004bb48, 0x1, 0x1)
	/usr/local/go/src/runtime/debug/stack.go:24 +0xa7
main.RecoveryMw.func1.1(0xa04fe0, 0xc4201a6000)
	/home/chinmaya/go/src/github.com/chinmaya1/gophercises/recover/main.go:66 +0xac
panic(0x833ca0, 0x9fc2e0)
	/usr/local/go/src/runtime/panic.go:502 +0x229
main.funcThatPanic()
	/home/chinmaya/go/src/github.com/chinmaya1/gophercises/recover/main.go:81 +0x39
main.PanicHandler(0xa04fe0, 0xc4201a6000, 0xc420132000)
	/home/chinmaya/go/src/github.com/chinmaya1/gophercises/recover/main.go:77 +0x20
net/http.HandlerFunc.ServeHTTP(0x930350, 0xa04fe0, 0xc4201a6000, 0xc420132000)
	/usr/local/go/src/net/http/server.go:1947 +0x44
net/http.(*ServeMux).ServeHTTP(0xc420487e00, 0xa04fe0, 0xc4201a6000, 0xc420132000)
	/usr/local/go/src/net/http/server.go:2337 +0x130
main.RecoveryMw.func1(0xa04fe0, 0xc4201a6000, 0xc420132000)
	/home/chinmaya/go/src/github.com/chinmaya1/gophercises/recover/main.go:72 +0x95
net/http.HandlerFunc.ServeHTTP(0xc4204ca9e0, 0xa04fe0, 0xc4201a6000, 0xc420132000)
	/usr/local/go/src/net/http/server.go:1947 +0x44
net/http.serverHandler.ServeHTTP(0xc4204cea90, 0xa04fe0, 0xc4201a6000, 0xc420132000)
	/usr/local/go/src/net/http/server.go:2694 +0xbc
net/http.(*conn).serve(0xc420080140, 0xa05360, 0xc420060100)
	/usr/local/go/src/net/http/server.go:1830 +0x651
created by net/http.(*Server).Serve
	/usr/local/go/src/net/http/server.go:2795 +0x27b`

	CreateLinks(str)
}

func TestPanic(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000/panic", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	PanicHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}

func TestM(t *testing.T) {
	tmp := listenAndServe
	defer func() {
		listenAndServe = tmp
	}()
	listenAndServe = func(addr string, handler http.Handler) error {
		return nil
	}
	assert.NotPanicsf(t, main, "they should be equal")
}
