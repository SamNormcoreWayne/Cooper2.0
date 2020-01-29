package errors

import (
    "fmt"
    "errors"
)

func unexpectedMethodError(wrong string, right string) (err error){
    errStr := fmt.Sprintf("Expeceted %s method but found %s", right, wrong)
    err = errors.New(errStr)
    return
}