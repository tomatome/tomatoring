package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", adminHandler)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func adminHandler(res http.ResponseWriter, req *http.Request) {
    if req.URL.Path == "/hello" {
       fmt.Fprintln(res, "hello, world")
       return
    }
    
     http.Redirect(res, req, "/index", http.StatusFound)
    
}
