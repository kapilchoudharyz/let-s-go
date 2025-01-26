package main

import (
	"auth-service/config"
	"auth-service/internal/auth"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("Request received: %v", r)
		log.Printf("Request type: %v", r.Method)
		log.Printf("http request type: %v", http.MethodGet)
		response := map[string]string{"status": "UP"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error writing response: %v", err)

		}
		// After last write to the response w the response is sent to the client.

	})
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/signup", auth.SignupHandler)
	log.Println("Server is starting on port " + cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
