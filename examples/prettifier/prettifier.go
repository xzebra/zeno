package main

import (
	"net/http"
	"strings"
	"text/template"

	"zeno"
)

// Demo website capable of display "pretty" math expressions
// acording to url path.
//
// ex: localhost:8080/3+2/5

func handleRequest(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Path
	expression = strings.TrimPrefix(expression, "/")
	tree, err := zeno.ToTree(expression)
	if err != nil {
		panic(err)
	}
	temp, err := template.ParseFiles("public/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, struct{ Content string }{
		Content: "$$" + tree.LaTeX() + "$$",
	})
}

func main() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
