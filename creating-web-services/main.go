package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ProductId      int    `json:"productId"`
	Manufacturer   string `json: "manufacturer"`
	Sku            string `json: "sku"`
	Upc            string `json: "upc"`
	PricePerUnit   string `json: "pricePerUnit"`
	QuantityOnHand int    `json: "quantityOnHand"`
	ProductName    string `json: productName`
}

var productList []Product

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(pJson)

	case http.MethodPost:
		var newProduct Product
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(body, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		if newProduct.ProductId != 0 {
			w.WriteHeader(http.StatusBadRequest)
		}
		newProduct.ProductId = getNextId()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)

	}
}
func productHandler(w http.ResponseWriter, r *http.Request){
	urlPath := strings.Split(r.URL.Path, "products/")
	productId, err := strconv.Atoi(urlPath[len(urlPath) -1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	product, id := findByProductId(productId)

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	



}

func init() {
	pJson := `[
		{
		  "productId": 1,
		  "manufacturer": "Johns-Jenkins",
		  "sku": "p5z343vdS",
		  "upc": "939581000000",
		  "pricePerUnit": "497.45",
		  "quantityOnHand": 9703,
		  "productName": "sticky note"
		},
		{
		  "productId": 2,
		  "manufacturer": "Hessel, Schimmel and Feeney",
		  "sku": "i7v300kmx",
		  "upc": "740979000000",
		  "pricePerUnit": "282.29",
		  "quantityOnHand": 9217,
		  "productName": "leg warmers"
		},
		{
		  "productId": 3,
		  "manufacturer": "Swaniawski, Bartoletti and Bruen",
		  "sku": "q0L657ys7",
		  "upc": "111730000000",
		  "pricePerUnit": "436.26",
		  "quantityOnHand": 5905,
		  "productName": "lamp shade"
		}]`

	err := json.Unmarshal([]byte(pJson), &productList)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productHandler)

	http.ListenAndServe(":5000", nil)

	// JSONTest()
}


// utils
func getNextId() int {
	return len(productList) + 1
}
func findByProductId(id int) (*Product, int){
	for i, p := range productList {
		if p.ProductId == id {
			return &p, i
		}
	}
	return nil, 0
}