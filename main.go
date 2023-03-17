package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerMap := map[string]http.HandlerFunc {
		http.MethodGet: handleGet,
		http.MethodPost: handlePost,
	}

	http.HandleFunc("/", listHandlersMap(handlerMap))

	http.ListenAndServe(":8001", nil)
}

func listHandlersMap(hm map[string]http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler, ok := hm[r.Method]
			if ok {
				handler(w, r)
				return
			}

			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Received GET request")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Received POST request")
}
