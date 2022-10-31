package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	enders "github.com/5HeadGames/web3-metadata-api/endersapi"
)

func buildReqest(encodedQuery *bytes.Buffer, graphApi string) *http.Request {
	toRequest, _ := http.NewRequest(
		"POST",
		graphApi,
		encodedQuery)
	toRequest.Header.Set("Content-Type", "application/json")
	return toRequest
}

func main() {
	apiKey := enders.GetEnvVars("SUBGRAPH_API")

	http.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {

		rw.Write([]byte("Enders Subgraph API"))
	})

	http.HandleFunc("/packs", func(rw http.ResponseWriter, req *http.Request) {
		address := req.URL.Query().Get("contractAddress")
		toQuery := enders.PacksQuery(address)
		formattedRequest := buildReqest(enders.EncodeQuery(toQuery), apiKey)

		client := &http.Client{Timeout: time.Second * 5}
		response, err := client.Do(formattedRequest)

		if err != nil {
			fmt.Printf("Error processing request %v", err)
		}

		responseData, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("Error reading data response %v", err)
		}

		rw.Write(responseData)
	})

	http.HandleFunc("/sales", func(rw http.ResponseWriter, req *http.Request) {
		address := req.URL.Query().Get("sellerAddress")
		toQuery := enders.SalesQuery(address)
		formattedRequest := buildReqest(enders.EncodeQuery(toQuery), apiKey)

		client := &http.Client{Timeout: time.Second * 5}
		response, err := client.Do(formattedRequest)

		if err != nil {
			fmt.Printf("Error processing request %v", err)
		}

		responseData, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("Error reading data response %v", err)
		}

		rw.Write(responseData)
	})

	http.HandleFunc("/buys", func(rw http.ResponseWriter, req *http.Request) {
		address := req.URL.Query().Get("buyerAddress")
		toQuery := enders.BuysQuery(address)
		formattedRequest := buildReqest(enders.EncodeQuery(toQuery), apiKey)

		client := &http.Client{Timeout: time.Second * 5}
		response, err := client.Do(formattedRequest)

		if err != nil {
			fmt.Printf("Error processing request %v", err)
		}

		responseData, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("Error reading data response %v", err)
		}

		rw.Write(responseData)
	})

	fmt.Println("Server listening on port 8080")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)
}
