package middleware

import (
    "log"
    "net/http"
    "mainapp/app/middleware/errors/"
)

func getParser(r *http.Request, keyNames []string) (keys []string, ok bool) {
    if r.Method != "GET" {
        /*
         * Raise an error here
         */
    }
    for _, key := range keyNames {
        keys, ok = r.URL.Query()[key]
        if !ok || len(keys[0]) < 1 {
            // We expected URL has a query here
            log.Println("URL Query param missing")
            return errors.getNoQueryError.getErr()
        }
    }
    return 
}