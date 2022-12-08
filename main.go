package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(r.Body)
		}
		bodyString := string(bodyBytes)
		fmt.Println("---------------------")
		fmt.Println(r.Header)
		fmt.Println(bodyString)
		fmt.Println("---------------------")
	})
	http.ListenAndServe(":8080", nil)
}
