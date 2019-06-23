package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	println(expression)
	tree, err := zeno.ToTree(expression)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	temp, err := template.ParseFiles("public/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := strconv.FormatFloat(tree.Operate(), 'f', -1, 64)

	temp.Execute(w, struct{ Expression, Result string }{
		Expression: tree.LaTeX(),
		Result:     result,
	})
}

func main() {
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})
	http.HandleFunc("/", handleRequest)
	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
