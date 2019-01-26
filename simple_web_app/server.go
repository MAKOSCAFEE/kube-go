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
	Category string `json:"category"`
}

func main() {
	port := 8080

	http.HandleFunc("/api/v1", handleSimpleWeb)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func handleSimpleWeb(w http.ResponseWriter, r *http.Request) {
	response := apiMessageResponse{Message: "Success", Status: 200}

	data, err := json.Marshal(response)

	if err != nil {
		panic("Ooops, There is problem here")
	}

	fmt.Fprint(w, string(data))
}
