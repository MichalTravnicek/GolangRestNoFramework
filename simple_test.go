package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "net/http"
    // "io"
    "encoding/json"
    // "fmt"
    // "time"
    "log"

    "golang-rest-no-framework/models"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	log.Println("Hello test")
}

func TestResponse(t *testing.T) {
    go main()
    // time.Sleep(100 * time.Millisecond)
    log.Println("Test http")
    // request, err := http.NewRequest("GET","https://cuddly-dollop-q795vjjq769v249px-8080.app.github.dev/products", nil)
    request, err := http.NewRequest("GET","http://localhost:8080/products/002", nil)
    res, err := http.DefaultClient.Do(request)
    log.Println("Response", res.Status)
    // resBody, err := io.ReadAll(res.Body)
    // log.Println(len(resBody))
    var response models.JsonResponse
    // json.Unmarshal(resBody, &product)
    defer res.Body.Close()
    json.NewDecoder(res.Body).Decode(&response)
    // product.Name = "name"
    //bodyReader := bytes.NewReader(resBody)
    log.Println("Response:", response.Message)
    log.Println("Product:", response.Data)
    log.Println("response", res)
    log.Println("Error", err)
    quit <- true
    assert.Equal(t,response.Message,"Get Product Success","should be equal")
}