package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"golang-rest-no-framework/models"
)

// temp db
var (
	quit     = make(chan bool)
	database = make(map[string]models.Product)
)

func setJsonResp(message []byte, httpCode int, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

func unmarshallingError(res http.ResponseWriter) {
	message := []byte(`{"message": "Error unmarshalling data"}`)
	setJsonResp(message, http.StatusBadRequest, res)
}

func products(res http.ResponseWriter, req *http.Request) {
	// check http method GET

	if req.Method == "GET" {
		var products []models.Product

		// loop through db
		for _, product := range database {
			products = append(products, product)
		}

		// Create a JSON response struct
		response := models.JsonResponse{
			Message: "Get All Product Success",
			Data:    products,
		}

		productJson, err := json.Marshal(response)

		if err != nil {
			message := []byte(`{"message": "Error marshalling data"}`)
			setJsonResp(message, http.StatusInternalServerError, res)
			return
		}

		// set response
		setJsonResp(productJson, http.StatusOK, res)
		return
	}

	if req.Method == "POST" {

		var product models.Product

		payload := req.Body

		defer req.Body.Close()
		// add product
		err := json.NewDecoder(payload).Decode(&product)

		if err != nil {
			unmarshallingError(res)
			return
		}

		// add product to db
		database[product.ID] = product

		message := []byte(`{"message": "Add new product success"}`)

		setJsonResp(message, http.StatusCreated, res)
	}

}

func productById(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" && req.Method != "DELETE" && req.Method != "PUT" {
		message := []byte(`{"message": "Invalid HTTP Method"}`)
		setJsonResp(message, http.StatusMethodNotAllowed, res)
		return
	}

	// // using inline if
	// if _, ok := req.URL.Query()["id"]; !ok {
	// 	message := []byte(`{"message": "Please provide id"}`)
	// 	setJsonResp(message, http.StatusBadRequest, res)
	// 	return
	// }

	// id := req.URL.Query().Get("id")

	id := strings.TrimPrefix(req.URL.Path, "/products/")

	log.Println("Requested product:", id)

	// product, boolean
	productData, ok := database[id]

	if !ok {
		message := []byte(`{"message": "Product not found"}`)
		setJsonResp(message, http.StatusNotFound, res)
		return
	}

	if req.Method == "DELETE" {
		delete(database, id)
		message := []byte(`{"message": "Product deleted"}`)
		setJsonResp(message, http.StatusNoContent, res)
		return
	}

	if req.Method == "PUT" {

		// get payload
		payload := req.Body

		defer req.Body.Close()

		// UPDATE PRODUCT
		var updatedProduct models.Product

		// add product

		err := json.NewDecoder(payload).Decode(&updatedProduct)

		if err != nil {
			message := []byte(`{"message": "Error unmarshalling data"}`)
			setJsonResp(message, http.StatusInternalServerError, res)
			return
		}

		// update product to db
		database[updatedProduct.ID] = updatedProduct

		message := []byte(`{"message": "Update product success"}`)

		setJsonResp(message, http.StatusNoContent, res)
		return
	}

	// Create a JSON response struct
	response := models.JsonResponse{
		Message: "Get Product Success",
		Data:    []models.Product{productData},
	}

	productJson, err := json.Marshal(response)

	if err != nil {
		unmarshallingError(res)
		return
	}
	setJsonResp(productJson, http.StatusOK, res)

}

var user string
var password string
var db string
var host string
var port string
var ssl string

func init() {
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	db = os.Getenv("POSTGRES_DB")
	host = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("POSTGRES_PORT")
	ssl = os.Getenv("POSTGRES_SSL")
}

func main() {
	// dsn := "host=db user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, db, port, ssl)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	log.Println("Connection:")
	log.Println(db.DB())
	log.Println("Errors:", err)

	log.Println("Application started")

	// init db
	database["001"] = models.Product{ID: "001", Name: "Pisang Goreng", Price: 10.99, Quantity: 10}
	database["002"] = models.Product{ID: "002", Name: "Teh Botol", Price: 5.99, Quantity: 20}

	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	message := []byte(`{"message": "Server up and running"}`)
	// 	setJsonResp(message, http.StatusOK, res)

	// })

	http.HandleFunc("/products", products)
	http.HandleFunc("/products/", productById)

	go func() {
		err := http.ListenAndServe(":8080", nil)

		log.Println("Server started")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	func() {
		for {
			select {
			case <-quit:
				log.Println("Quitting")
				return
			default:
				// log.Println("Sleeping")
				time.Sleep(100 * time.Millisecond)
				// Do other stuff
			}
		}
	}()

	//time.Sleep(1000 * time.Millisecond)
}
