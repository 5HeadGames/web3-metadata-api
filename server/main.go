package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

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

func requestData(rawQuery map[string]string, endpoint string) []byte {
	newRequest := buildReqest(enders.EncodeQuery(rawQuery), endpoint)

	client := &http.Client{Timeout: time.Second * 5}
	response, err := client.Do(newRequest)

	if err != nil {
		fmt.Printf("Error processing request %v", err)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error reading data response %v", err)
	}

	return responseData
}

func main() {
	apiKey := enders.GetEnvVars("SUBGRAPH_API")
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {

		c.IndentedJSON(http.StatusOK, map[string]string{
			"message": "Enders Subgraph API",
		})
	})

	route.GET("/packs", func(c *gin.Context) {
		address := c.Query("contractAddress")
		toQuery := enders.PacksQuery(address)

		rawResult := requestData(toQuery, apiKey)

		decodedResult := enders.PacksResponse{}
		jsonErr := json.Unmarshal(rawResult, &decodedResult)

		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		c.IndentedJSON(http.StatusOK, decodedResult)
	})

	route.GET("/sales", func(c *gin.Context) {
		address := c.Query("sellerAddress")
		toQuery := enders.SalesQuery(address)

		rawResult := requestData(toQuery, apiKey)

		decodedResult := enders.SalesResponse{}

		jsonErr := json.Unmarshal(rawResult, &decodedResult)

		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		c.IndentedJSON(http.StatusOK, decodedResult)
	})

	route.GET("/buys", func(c *gin.Context) {
		address := c.Query("buyerAddress")
		toQuery := enders.BuysQuery(address)

		rawResult := requestData(toQuery, apiKey)

		decodedResult := enders.BuysResponse{}

		jsonErr := json.Unmarshal(rawResult, &decodedResult)

		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		c.IndentedJSON(http.StatusOK, decodedResult)
	})

	fmt.Println("Server listening on port 8080")
	log.Panic(
		route.Run("localhost:8080"),
	)
}
