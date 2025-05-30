package main

import (
    "log"
    "net/http"

    "github.com/noahsignt/blackout/be/handler"
)

func main() {
    r := handler.NewRouter()

    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
