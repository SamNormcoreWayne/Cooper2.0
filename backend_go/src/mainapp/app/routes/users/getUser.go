package users

import (
    "fmt"
    "net/http"
    "net/url"
    "encoding/json"
    "mainapp/app/middleware"
)

// DB_BASE_URL Database Address
const DB_BASE_URL string = "http://cooper-database-api:8080"
// Type is an HTTP content-type key
const Type string = "Content-Type"
// contentT is an HTTP content-type value
const contentT string = "application/json"

func getUser(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r)
    /**
     * TODO: Encapsulate to getParser(). Done
     */
    keys, _, parserErr := middleware.GetParser(r, []string{"id"})

    if (parserErr != nil) {
        /**
        * Paser error handling
        */
        w.Header().Set(Type, contentT)
        w.WriteHeader(http.StatusBadRequest)
        /*
         * Need a middleware to handle this error.
         * Need a better design. Functions might be duplicated here.
         */
    }

    userID := keys[0]
    dbURL, _ := url.Parse(DB_BASE_URL)
    dbURL.Query().Set("users", userID)
    res, err := http.Get(dbURL.String()) //~TODO: make this line async~, concurrency shall not be resolved here, but it should be resolved when the req reaches our server.
    if err == nil {
        w.Header().Set(Type, contentT)
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(res.Body)
        /*
         * Check here in testing. I am not sure 
         * whether res.Body could be parsered properly
         */
    } else {
        w.Header().Set(Type, contentT)
        w.WriteHeader(http.StatusBadRequest)
    }
    defer res.Body.Close()
}

// GetUser exports getUser()
var GetUser = getUser