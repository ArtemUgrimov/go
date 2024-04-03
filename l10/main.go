package main

import (
	"fmt"
	"net/http"
)

func main() {
	tm := NewTaskManager()
	router := tm.getRouter()
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print("Server run error", err)
	}
}
