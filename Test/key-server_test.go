package main

import (
    "flag"
    "fmt"
    "net/http"
    "net/http/httptest"
    "os"
    "strconv"
    "testing"
)

var (
    maxSize int
)

func init() {
    // Register the maxSize flag to capture dynamic values
    flag.IntVar(&maxSize, "max-size", 1024, "maximum key size for test cases")
}

func TestMain(m *testing.M) {
    // Parse the flags to initialize maxSize from command-line arguments
    flag.Parse()
    os.Exit(m.Run())
}

func TestKeyHandlerValidWithDynamicMaxSize(t *testing.T) {
    // Use a valid key length less than or equal to maxSize
    validLength := maxSize - 1

    req, err := http.NewRequest("GET", fmt.Sprintf("/key?length=%d", validLength), nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(keyHandler)
    handler.ServeHTTP(rr, req)

    // Check the status code is 200 OK
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status code 200, got %d", status)
    }

    // Check the response body length (should match the validLength)
    if len(rr.Body.Bytes()) != validLength {
        t.Errorf("Expected %d bytes, got %d bytes", validLength, len(rr.Body.Bytes()))
    }
}

func TestKeyHandlerInvalidWithDynamicMaxSize(t *testing.T) {
    // Use a key length greater than maxSize to trigger invalid request
    invalidLength := maxSize + 1

    req, err := http.NewRequest("GET", fmt.Sprintf("/key?length=%d", invalidLength), nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(keyHandler)
    handler.ServeHTTP(rr, req)

    // Check that the response status is 400 Bad Request
    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("Expected status code 400, got %d", status)
    }

    // Check the error message in the response body
    expectedErrorMsg := "Invalid key length or exceeds max size\n"
    if rr.Body.String() != expectedErrorMsg {
        t.Errorf("Expected '%s', got '%s'", expectedErrorMsg, rr.Body.String())
    }
}