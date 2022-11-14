package endersapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func EncodeQuery(baseQuery map[string]string) *bytes.Buffer {
	jsonResult, err := json.Marshal(baseQuery)

	if err != nil {
		fmt.Printf("Encoding query failed %v", err)
	}

	return bytes.NewBuffer(jsonResult)
}

func PacksQuery(tokenId string) map[string]string {
	var payload = map[string]string{
		"query": fmt.Sprintf("query { erc1155Tokens ( where : { contract : \"%s\" } ) { totalSupply { value } } }", tokenId),
	}

	return payload
}

func SalesQuery(sellerAddress string) map[string]string {
	var payload = map[string]string{
		"query": fmt.Sprintf(
			"query {sales (where : {seller : \"%s\" }, orderBy : timestamp, orderDirection: desc) {id, amount, duration, seller{id}, nft{id, contract {id}}, price, timestamp, status}}",
			sellerAddress),
	}

	return payload
}

func AllSalesQuery() map[string]string {
	var payload = map[string]string{
		"query": `query {sales (orderBy : timestamp, orderDirection: desc)
			{id, amount, duration, seller{id}, nft{id, contract {id}},
			price, timestamp, status}}`,
	}
	return payload
}

func BuysQuery(buyerAddress string) map[string]string {
	var payload = map[string]string{
		"query": fmt.Sprintf("query {buys ( where : { buyer : \"%s\" } ) {buyer{id}, sale{id, seller{id}, nft{id, contract {id}}}, timestamp}}", buyerAddress),
	}

	return payload
}
