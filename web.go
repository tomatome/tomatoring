package main

import (
    "fmt"
    "net/http"
    "html/template"
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
    
    t, err := template.ParseFiles("index.html")
    if (err != nil) {
        fmt.Println(err)
    }
    t.Execute(res, nil)
    
}
