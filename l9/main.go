package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// render html

func main() {
	http.HandleFunc("/", emptyPathHandler)

	http.ListenAndServe(":8000", nil)
}

func emptyPathHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		fmt.Fprint(w, "Error reading 'a'")
		return
	}
	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		fmt.Fprint(w, "Error reading 'b'")
		return
	}
	op := r.URL.Query().Get("operation")

	if op == "+" {
		fmt.Fprintf(w, "sum = %d", a+b)
	} else if op == "-" {
		fmt.Fprintf(w, "dif = %d", a-b)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid operation")
	}
}
