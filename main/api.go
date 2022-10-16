package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/trades", func(rw http.ResponseWriter, req *http.Request) {
		address := req.URL.Query().Get("buyerAddress")
		rw.Write([]byte(fmt.Sprintf("Address, %s", address)))
	})

	http.HandleFunc("/packs", func(rw http.ResponseWriter, req *http.Request) {
		contract := req.URL.Query().Get("contractAddress")
		rw.Write([]byte(fmt.Sprintf("Contract Pack, %s", contract)))
	})

	http.HandleFunc("/sales", func(rw http.ResponseWriter, req *http.Request) {
		address := req.URL.Query().Get("userAddress")
		rw.Write([]byte(fmt.Sprintf("Address, %s", address)))
	})

	http.ListenAndServe(":8080", nil)
}
