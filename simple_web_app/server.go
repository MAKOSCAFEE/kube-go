/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-26 23:54:50
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-27 00:17:22
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type apiMessageResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type product struct {
	Name     string `json:"name"`
	ID       int    `json:"id,string"`
	Price    int    `json:"price,string"`
	Category string `json:"category"`
}

func main() {
	port := 8080

	http.HandleFunc("/api/v1", handleSimpleWeb)
	http.HandleFunc("/api/v1/products", handleProducts)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	response := product{Name: "Sugar", Price: 2300, ID: 3423, Category: "Food"}

	encoder := json.NewEncoder(w)

	encoder.Encode(&response)
}

func handleSimpleWeb(w http.ResponseWriter, r *http.Request) {
	response := apiMessageResponse{Message: "Success", Status: 200}

	encoder := json.NewEncoder(w)

	encoder.Encode(&response)
}
