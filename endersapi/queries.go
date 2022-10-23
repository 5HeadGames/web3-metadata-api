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

var PacksQuery = map[string]string{
	"query": `query GetSupply ($tokenId: ID!) {
        erc1155Tokens ( where : { contract : $tokenId } ) {
          totalSupply {
            value
          }
        }
    }`,
}
