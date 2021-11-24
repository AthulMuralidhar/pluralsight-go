package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		name := r.URL.Query()["name"][0]

		// var name string = names[0] 

		// just string
		// rw.Write([]byte("writting on the inteface " + name))

		// json encoder
		outputMap := map[string]string{"name": name}
		jsonEncoder := json.NewEncoder(rw)
		jsonEncoder.Encode(outputMap)
	})

	// nil here == default handler
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}