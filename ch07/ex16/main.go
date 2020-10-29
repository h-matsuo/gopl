package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/h-matsuo/gopl/ch07/ex14/eval"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	e := req.URL.Query().Get("expr")
	expr, err := eval.Parse(e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "expr not specified")
		return
	}

	env := eval.Env{}
	for query, value := range req.URL.Query() {
		if query == "expr" {
			continue
		}
		value, err := strconv.ParseFloat(value[0], 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s: %v", query, err)
			return
		}
		env[eval.Var(query)] = value
	}

	fmt.Fprintf(w, "%f\n", expr.Eval(env))
}
