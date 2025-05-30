package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Server is running!"))
    })

    log.Println("🚀 Server listening on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Server failed:", err)
    }
}
