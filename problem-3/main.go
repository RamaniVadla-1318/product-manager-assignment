package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("Monday, 02-Jan-2006 15:04:05 MST")
    fmt.Fprintf(w, "<h1>Current Date and Time</h1><p>%s</p>", currentTime)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on :8080...")
    http.ListenAndServe(":8080", nil)
}

