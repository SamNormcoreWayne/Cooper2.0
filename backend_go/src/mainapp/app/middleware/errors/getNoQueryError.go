package errors

import (
    "errors"
    "fmt"
)

func getErr() {
    getError := errors.New("GET: No queries")
    fmt.Println(getError)
}