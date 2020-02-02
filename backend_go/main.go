package main

import (
    "fmt"
    "net/http"
    "log"
    routers "mainapp/app/routes"
)

func main() {
    fmt.Println("Hello, world")
    fmt.Println("This is an app for Cooper2.0/Backend")

    fmt.Println("Starting server at port 9000")

    http.HandleFunc("/", routers.Router)
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal("ListenAndServer: ", err)
    }
}
