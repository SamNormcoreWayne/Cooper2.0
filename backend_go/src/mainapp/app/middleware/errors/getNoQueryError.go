package errors

import (
    "errors"
    "fmt"
)

func getNoQueryError() (err error){
    err = errors.New("GET: No queries")
    fmt.Println(err)
    return
}