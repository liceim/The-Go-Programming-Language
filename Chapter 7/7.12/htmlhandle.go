/*
Exercise 7.12: Change the handler for/listto print its output as an HTML table, not text. You may find thehtml/templatepackage (§4.6) useful.

*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/createUpdate", db.createUpdate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

//!+template
var itemTable = template.Must(template.New("Items").Parse(`
<h1>Items</h1>
<table>
    <tr>
        <th> Item </th>
        <th> Price </th>
    </tr>
    {{ range $k, $v := . }}
        <tr>
            <td>{{ $k }}</td>
            <td>{{ $v }}</td>
        </tr>
    {{end}}
</table>
`))

//!-template

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	itemTable.Execute(w, db)
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

var mux sync.Mutex

// create and update
func (db database) createUpdate(w http.ResponseWriter, req *http.Request) {
	// syntax: /create?item=cpu?price=5
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item != "" && price != "" {
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "item: %q, price: %q\n", item, price)
			return
		}
		// create item
		mux.Lock()
		db[item] = dollars(p)
		mux.Unlock()
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprintf(w, "item: %q, price: %q\n", item, price)
}
