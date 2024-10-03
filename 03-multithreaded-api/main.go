package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloApi)
	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func HelloApi(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	w.Header().Add("Content-Type", "application/json")
	wg.Add(1)
	go func() {
		defer wg.Done()
		json.NewEncoder(w).Encode("hello world")
	}()
	wg.Wait()
}
