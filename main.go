package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("02-Jan-2006 15:04:05 MST")
    fmt.Fprintf(w, "Current Date and Time: %s", currentTime)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

