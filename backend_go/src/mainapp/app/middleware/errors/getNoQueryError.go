package errors

import (
    "errors"
    "fmt"
)

// GetNoQueryError handle a error if there are no queries input
func GetNoQueryError() (err error){
    err = errors.New("GET: No queries")
    fmt.Println(err)
    return
}