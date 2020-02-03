package users_test

import (
    "log"
    "testing"
    "net/http/httptest"
    "net/http"
    "github.com/jarcoal/httpmock"
    "mainapp/app/routes"
    "io/ioutil"
)

type getUserRes struct {
    ID string `json:"userId"`
}

const DB_BASE_URL string = "http://cooper-database-api:8080"

func TestGetUser(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    // Mock DB
    userRes := &getUserRes{ID: "4"}
    httpmock.RegisterResponder("GET", DB_BASE_URL + "/users/4", 
        func (req *http.Request) (*http.Response, error) {
            resp, err := httpmock.NewJsonResponse(200, userRes)
            if (err != nil) {
                return httpmock.NewStringResponse(500, "User does not exist"), nil
            }
            return resp, nil
        })
    
    // Test no query error
    testHandler := http.HandlerFunc(routers.GetUser)
    testReq := httptest.NewRequest("GET", "/api/user/getUser", nil)
    testRes := httptest.NewRecorder()
    testHandler.ServeHTTP(testRes, testReq)
    res := testRes.Result()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    expectedBody := "GET /api/user/getUser StatusBadRequest"
    if string(body) ==  expectedBody {
        t.Logf("Should fail to get user due to no query")
    }
}