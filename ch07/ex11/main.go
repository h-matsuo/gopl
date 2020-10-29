package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
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

var wg sync.WaitGroup

func (db database) update(w http.ResponseWriter, req *http.Request) {
	wg.Wait()
	wg.Add(1)
	defer wg.Done()
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid price")
		return
	}
	if _, ok := db[item]; ok {
		db[item] = dollars(price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	db.list(w, req)
}
