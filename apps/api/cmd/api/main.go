package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	fmt.Println("API running on port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
