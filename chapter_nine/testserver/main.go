package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", sleepHandler)
	http.ListenAndServe(":8080", nil)
}

func sleepHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintf(w, "Server slept for 0.1 seconds")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server slept for 0.1 seconds"))
}
