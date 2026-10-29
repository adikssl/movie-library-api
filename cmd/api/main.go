package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method allows GET", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Movie Library API is running")
}

func main() {
	http.HandleFunc("/health", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Critical server startup error:", err)
	}

}
