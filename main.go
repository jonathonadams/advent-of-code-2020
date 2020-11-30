package main

import (  
    "fmt"
    "net/http"
)

func main() { 
    fmt.Println("Server is running on port :9090")

    err := http.ListenAndServe(":9090", http.FileServer(http.Dir("./public")))
    
    if err != nil {
        fmt.Println("Failed to start server", err)
        return
    }
}