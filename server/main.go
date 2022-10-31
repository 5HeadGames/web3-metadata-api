package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func makeReqest(encodedQuery *bytes.Buffer, apiKey string) *http.Request {
	toRequest, _ := http.NewRequest(
		"POST",
		apiKey,
		encodedQuery)
	toRequest.Header.Set("Content-Type", "application/json")
	return toRequest
}

func main() {
	// apiKey := enders.GetEnvVars("SUBGRAPH_API")

	http.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {

		rw.Write([]byte("Enders Subgraph API"))

		// client := &http.Client{Timeout: time.Second * 5}
		// response, err := client.Do(newRequest)
		// if err != nil {
		// 	fmt.Printf("Error processing request %v", err)
		// }

		// responseData, err := ioutil.ReadAll(response.Body)

		// if err != nil {
		// 	fmt.Printf("Error reading data response %v", err)
		// }

		// fmt.Println(string(responseData))
	})

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

	fmt.Println("Server listening on port 8080")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)
}
