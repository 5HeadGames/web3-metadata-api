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
