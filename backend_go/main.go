package main

import (
    "fmt"
    "net/http"
    "log"
    "io"
    "mainapp/app/routes"
)

func main() {
    fmt.Println("Hello, world")
    fmt.Println("This is an app for Cooper2.0/Backend")

    fmt.Println("Starting server at port 9000")

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello World!")
    })
    http.HandleFunc("/home", routers.HelloPage)
    http.HandleFunc("/api/user/getUser", routers.GetUser)
    log.Fatal(http.ListenAndServe(":9000", nil))
}
