package middleware

import (
    "log"
    "net/http"
)

func getParser(r *http.Request) (keys []string, ok bool) {
    if r.Method != "GET" {
        /*
         * Raise an error here
         */
    }
    keys, ok = r.URL.Query()["id"]
    if !ok || len(keys[0]) < 1 {
        log.Println("URL Query param missing")
        //Raise an error here!
    }
    return
}