package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var listenAndServe = http.ListenAndServe

//Handler function will handle all the routes
//GetHandler function is using MUX to handle the routes
func GetHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", SourceCodeHandler)
	mux.HandleFunc("/panic", PanicHandler)

	return mux
}

func main() {
	listenAndServe(":8000", RecoveryMw(GetHandler()))
}

//SourceCodeHandler function is used to debug the errors.
//SourceCodeHandler function to handle source code of given file
func SourceCodeHandler(w http.ResponseWriter, r *http.Request) {

	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	io.Copy(b, file)
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, b.String())
	style := styles.Get("github")
	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(), html.LineNumbersInTable(), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)

}

//RecoveryMw to recover panics from program
func RecoveryMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, CreateLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

//PanicHandler to handle panic function
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	funcThatPanic()
}

//PanicFunction to display panic
func funcThatPanic() {
	panic("oh no!")
}

// CreateLinks takes the string and form the links at the line numbers
func CreateLinks(stack string) string {
	re := regexp.MustCompile("\t.*:[0-9]*")
	lines := re.FindAllString(stack, -1)

	re = regexp.MustCompile(":")
	for _, line := range lines {
		splits := re.Split(line, -1)
		link := "<a href='/debug?path=" + splits[0] + "&line=" + splits[1] + "'>" + line + "</a>"
		reg := regexp.MustCompile(line)
		stack = reg.ReplaceAllString(stack, link)
	}

	return stack
}
