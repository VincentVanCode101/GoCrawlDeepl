package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("coverage/coverage.html")
		if err != nil {

			log.Printf("Error reading HTML file: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		_, err = w.Write(data)

		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	})

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
