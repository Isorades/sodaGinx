package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Message struct {
    Message string `json:"message"`
}

type Response struct {
    Response string `json:"response"`
}

func main() {
    http.HandleFunc("/", handleRequest)
    
    port := 8082
    fmt.Printf("Server 2 starting on port %d...\n", port)
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        log.Fatal(err)
    }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var msg Message
    if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response := Response{
        Response: fmt.Sprintf("Server 8082 收到消息: %s", msg.Message),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
} 