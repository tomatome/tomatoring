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
    if r.URL.Path == "/hello" {
       fmt.Fprintln(res, "hello, world")
       return
    }
    
     http.Redirect(w, r, "/index", http.StatusFound)
    
}
