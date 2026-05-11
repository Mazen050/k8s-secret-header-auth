package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    expectedKey := os.Getenv("API_KEY")
    if expectedKey == "" {
        log.Fatal("API_KEY env var not set")
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        incoming := r.Header.Get("X-Api-Key")
        if incoming != expectedKey {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        fmt.Fprintf(w, "Authorized! Header validated against Kubernetes Secret.\n")
    })

    log.Println("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}