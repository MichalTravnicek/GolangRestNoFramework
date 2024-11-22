package main

import (
    "testing"
    "net/http"
    "time"
    "log"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	log.Println("Hello test")
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    go main()
    time.Sleep(100 * time.Millisecond)
    log.Println("Test http")
    request, err := http.NewRequest("GET","/", nil)
    // log.Println("Request", request.Response.Status)
    log.Println("Request", request.Body)
    log.Println("Error", err)
    quit <- true
}