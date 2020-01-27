package routers

import (
    "net/http"
    "io"
    "mainapp/app/routes/users"
)

func router() {
    http.HandleFunc("/home", helloPage)
    http.HandleFunc("/api/user/getUser", users.getUser)
}

func helloPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world! \nThis is the gateway of Mr.Cooper2.0")
}