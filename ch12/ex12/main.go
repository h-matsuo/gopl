package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/h-matsuo/gopl/ch12/ex12/params"
)

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		UUID   string   `http:"uid" validate:"uuid"`
		Emails []string `http:"e" validate:"email"`
	}
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
