package users

import (
    "fmt"
    "io"
    "net/http"
    "net/url"
    "mainapp/app/middleware"
    "io/ioutil"
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
    w.Header().Set(Type, contentT)
    /**
     * TODO: Encapsulate to getParser(). Done
     */
    keys, _, parserErr := middleware.GetParser(r, []string{"userId"})
    if (parserErr != nil) {
        /**
        * Paser error handling
        */
        w.WriteHeader(http.StatusBadRequest)
        /*
         * Need a middleware to handle this error.
         * Need a better design. Functions might be duplicated here.
         */
         io.WriteString(w, "GET /api/user/getUser StatusBadRequest")
    } else {
        userID := (string)(keys[0])
        dbURL, _ := url.Parse(DB_BASE_URL + "/users/" + userID)
        res, err := http.Get(dbURL.String()) //~TODO: make this line async~, concurrency shall not be resolved here, but it should be resolved when the req reaches our server.
        if err == nil {
            w.WriteHeader(http.StatusOK)
            body, _ := ioutil.ReadAll(res.Body)
            fmt.Println(body)
            io.WriteString(w, (string)(body))
            /*
            * Check here in testing. I am not sure 
            * whether res.Body could be parsered properly
            */
            res.Body.Close()
        } else {
            w.WriteHeader(http.StatusBadRequest)
        }
    }
}

// GetUser exports getUser()
var GetUser = getUser