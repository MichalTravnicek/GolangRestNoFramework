package main

import (
    "testing"
    "net/http"
    // "io"
    "encoding/json"
    // "fmt"
    "time"
    "log"

    "github.com/raymondddenny/golang-rest-no-framework/models"
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
    time.Sleep(1000 * time.Millisecond)
    log.Println("Test http")
    // request, err := http.NewRequest("GET","https://cuddly-dollop-q795vjjq769v249px-8080.app.github.dev/products", nil)
    request, err := http.NewRequest("GET","http://localhost:8080/product/001", nil)
    res, err := http.DefaultClient.Do(request)
    log.Println("Response", res.Status)
    // resBody, err := io.ReadAll(res.Body)
    // log.Println(len(resBody))
    var product models.Product
    // json.Unmarshal(resBody, &product)
    json.NewDecoder(res.Body).Decode(&product)
    //bodyReader := bytes.NewReader(resBody)
    log.Println("Product", product)
    log.Println("response", res)
    log.Println("Error", err)
    quit <- true
}