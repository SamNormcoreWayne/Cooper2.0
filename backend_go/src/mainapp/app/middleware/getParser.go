package middleware

import (
    "log"
    "net/http"
    "mainapp/app/middleware/errors"
)

// GetParser to get param from a URL that method is GET
func GetParser(r *http.Request, keyNames []string) (keys []string, ok bool, err error) {
    if r.Method != "GET" {
        /*
         * Raise an error here
         */
        return nil, false, errors.UnexpectedMethodErr(r.Method, "GET")
    }
    for _, key := range keyNames {
        keys, ok = r.URL.Query()[key]
        log.Println(keys, ok)
        if !ok || len(keys) == 0 {
            // We expected URL has a query here
            log.Println("URL Query param missing")
            return nil, false, errors.GetNoQueryError()
        }
    }
    return  keys, ok, nil
}