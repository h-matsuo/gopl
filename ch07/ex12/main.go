package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>Items List</h1>
<table>
<tr>
  <th style='text-align: left'>Item</th>
  <th style='text-align: right'>Price</th>
</tr>
{{ range $item, $price := . }}
<tr>
  <td style='text-align: right'>{{ $item }}</td>
  <td style='text-align: right'>{{ $price }}</td>
</tr>
{{end}}
</table>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := issueList.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error\n%v", err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
