package errors

import (
    "errors"
    "fmt"
)

func getErr() (err error){
    err = errors.New("GET: No queries")
    fmt.Println(err)
    return
}