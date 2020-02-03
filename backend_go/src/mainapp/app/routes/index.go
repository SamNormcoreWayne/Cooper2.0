package routers

import (
    "net/http"
    "io"
    "mainapp/app/routes/users"
)

var GetUser = users.GetUser
var HelloPage = func (w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world! \nThis is the gateway of Mr.Cooper2.0")
}

// Router exports router()