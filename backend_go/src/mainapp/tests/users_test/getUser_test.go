package users_test

import (
    "testing"
    "net/http/httptest"
    "net/http"
    "fmt"
    "encoding/json"
    "github.com/jarcoal/httpmock"
    "mainapp/app/routes/users"
)

type getUserRes struct {
    ID string `json:"userId"`
}

const DB_BASE_URL string = "http://cooper-database-api:8080"

func TestGetUser(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    userRes := &getUserRes{ID: "4"}
    res, _ := json.Marshal(userRes)
    httpmock.RegisterResponder("GET", DB_BASE_URL + "/users/4", 
        func (req *http.Request) (*http.Response, error) {
            resp, err := httpmock.NewJsonResponse(200, userRes)
            if (err != nil) {
                return httpmock.NewStringResponse(500, "User does not exist"), nil
            }
            return resp, nil
        })

    testServer := httptest.NewServer(handler)
    defer testServer.close()
}