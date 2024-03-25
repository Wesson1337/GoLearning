package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"html/template"
)

type dollars float32
type database map[string]dollars

var itemMutex = sync.Mutex{}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, req *http.Request) {
	// html table of items and prices with html/template package
	tmpl := template.Must(template.New("items").Parse(`
		<h1>Items</h1>
		<table>
			<tr>
				<th>Item</th>
				<th>Price</th>
			</tr>
			{{range $item, $price := .}}
				<tr>
					<td>{{$item}}</td>
					<td>{{$price}}</td>
				</tr>
			{{end}}
		</table>
	`))
	tmpl.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing price\n")
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	intPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	} 

	if intPrice < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price must be non-negative: %q\n", price)
		return
	}

	dollarPrice := dollars(intPrice)

	itemMutex.Lock()
	db[item] = dollarPrice
	itemMutex.Unlock()

	fmt.Fprintf(w, "%s\n", dollarPrice)
}