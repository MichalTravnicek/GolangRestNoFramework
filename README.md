# Golang API without Framework

This repository contains a Golang API project that implements CRUD (Create, Read, Update, Delete) operations for products. It uses an array of maps as the database.

## Routes

1. **Check Server** (GET)
   - Endpoint: `localhost:8080`
   - Description: Check if the server is running.

2. **Get All Products** (GET)
   - Endpoint: `localhost:8080/products`
   - Description: Retrieve a list of all products.

3. **Get Product by ID** (GET)
   - Endpoint: `localhost:8080/product?id=003`
   - Description: Retrieve a product by its ID.

4. **Update Product by ID** (PUT)
   - Endpoint: `localhost:8080/product?id=003`
   - Description: Update a product's information by its ID.

5. **Delete Product by ID** (DELETE)
   - Endpoint: `localhost:8080/product?id=003`
   - Description: Delete a product by its ID.

6. **Add New Product** (POST)
   - Endpoint: `localhost:8080/products`
   - Description: Add a new product to the database.

## Database

The database for this API project is implemented using an array of maps. Each map represents a product with its attributes such as ID, name, description, and price. It provides a simple and lightweight way to store and manipulate data.

## Usage

1. Clone this repository to your local machine.

2. Ensure you have Go installed.

3. Navigate to the project directory.

4. Run the API server:
   ```shell
   go run main.go

You can now make requests to the specified endpoints using tools like *curl* or *Postman*.

## Contributing

Feel free to contribute to this project by submitting pull requests or reporting issues. Your contributions are highly appreciated!
