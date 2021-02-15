package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// 读取
func (d *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range *d {
		fmt.Fprintf(w, "%s :%s\n", item, price)
	}
}

// 创建
func (d *database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(p, 10)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error: price: %s\n", p)
		return
	}
	(*d)[item] = dollars(price)
}

// 更新
func (d *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(p, 10)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error: price: %s\n", p)
		return
	}
	if _, ok := (*d)[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error: no such item: %s\n", item)
		return
	}
	(*d)[item] = dollars(price)
}

// 删除
func (d *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := (*d)[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error: no such item: %s\n", item)
		return
	}
	delete(*d, item)
}

func main() {
	db := database{
		"shore": 33,
		"shirt": 2345,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/add", db.add)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
