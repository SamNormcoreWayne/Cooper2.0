package routers

import (
    "net/http"
    "io"
    "mainapp/app/routes/users"
)

func router(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "MainPage")
    http.HandleFunc("/home", helloPage)
    http.HandleFunc("/api/user/getUser", users.GetUser) //Resolve concurrency
}

func helloPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world! \nThis is the gateway of Mr.Cooper2.0")
}

// Router exports router()
var Router = router