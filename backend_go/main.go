package main

import (
    "fmt"
    "log"
    "net/http"
    "mainapp/app/routes"
)

func main() {
    fmt.Println("Hello, world")
    fmt.Println("This is an app for Cooper2.0/Backend")

    fmt.Println("Starting server at port 9000")

    http.HandleFunc("/", routers.index)
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal("ListenAndServr port: 9000 Error: ", err)
    }
}
