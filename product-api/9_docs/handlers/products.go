package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PacktPublishing/Building-Microservices-with-Go-Second-Edition/product-api/9_docs/data"
	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Products handler for getting and updating products
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts returns a new products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
func getProductID(r *http.Request) (int, error) {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	return strconv.Atoi(vars["id"])
}
