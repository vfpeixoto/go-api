package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "go-api/business"
)



func main() {
    r := mux.NewRouter()
    
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", foo)
    r.HandleFunc("/profiles", profiles)

    // Bind to a port and pass our router in
    http.ListenAndServe(":8000", r)
}

func foo(w http.ResponseWriter, r *http.Request){

  w.Header().Set("Server", "A Go Web Server")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)

  w.Write(business.GetInfo())

}

func profiles(w http.ResponseWriter, r *http.Request) {
  
    w.Header().Set("Server", "A Go Web Server")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)

    fmt.Printf(r.Method)

    switch r.Method {
        case "GET": 
            w.Write(business.GetUser())
        case "POST":
        case "PUT":
        case "DELETE":
        default:
            //snipped
    }
  
}