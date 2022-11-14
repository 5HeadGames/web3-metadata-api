# Enders API
Go REST API to serve metadata indexed from a subgraph and a database containing on-chain events and user generated data.

## Current Endpoints
### "/"
**Method**: GET
**Return Type**: { message: string }
**Response Example:** `{"message": "Enders API"}`

### "/packs"
**Method**: GET
**Query Params**: 
* contractAddress (string): Address for requested erc1155 token

**Return Type**: `{ data: [erc1155Tokens: []struct {
            totalSupply: struct {
                value: string
            }
        } ] 
    }
`

**Request Example**: `http://api-gateway-example/packs?contractAddress=0xad...ac`

**Response Example**: `{
    "data": {
        "erc1155Tokens": [
            {
                "totalSupply": {
                    "value": "0.000000000000000013"
                }
            },
            {
                "totalSupply": {
                    "value": "0.000000000000000014"
                }
            }, ...
        ]
    }
}`

### "/sales"
**Method**: GET
**Query Params**:
* OPTIONAL sellerAddress (string): The user address which created the sale transaction.

**Return Type**: `{ data: [ sales: []struct {
            ammount: string,
            duration: string,
            price: string,
            seller: struct {
                id: string
            },
            nft: struct {
                id: string,
                contract: struct {
                    id: string
                }
            },
            timestamp: string,
            status: string
        } ]
    }
`

**Request Example**: `http://api-gateway-example/sales?sellerAddress=0xb9...cd`

**Response Example**: `{
    "data": {
        "sales": [
            {
                "amount": "1942",
                "duration": "8640000000",
                "nft": {
                    "id": "0xb90dc9e354001e6260de670edd6abadb890c6ac9/0x3",
                    "contract": {
                        "id": "0xb90dc9e354001e6260de670edd6abadb890c6ac9"
                    }
                },
                "price": "4836642400000000000000",
                "seller": {
                    "id": "0xb9602f2442da97651d5f7e0435a4733b1a1145cd"
                },
                "timestamp": "1653696067",
                "status": "0"
            },
            {
                "amount": "3973",
                "duration": "8640000000",
                "nft": {
                    "id": "0xb90dc9e354001e6260de670edd6abadb890c6ac9/0x2",
                    "contract": {
                        "id": "0xb90dc9e354001e6260de670edd6abadb890c6ac9"
                    }
                },
                "price": "2418964600000000000000",
                "seller": {
                    "id": "0xb9602f2442da97651d5f7e0435a4733b1a1145cd"
                },
                "timestamp": "1653696028",
                "status": "0"
            }, ...
        ]
    }
}`

### "/buys"
**Method**: GET
**Query Params**:
* buyerAddress: The user address which initiated the buy transaction.

**Return Type**: ` { data : [ buys: struct {
            buyer: struct {
                id: string
            },
            sale : struct {
                id: string,
                seller: struct {
                    id: string
                }
                nft : struct {
                    id: string,
                    contract: struct {
                        id: string
                    }
                }
            },
            timestamp: string
        } ] 
    }
`

**Request Example**: `http://api-gateway-example/buys?buyerAddress=0x35....1e`

**Response Example**: `{
    "data": {
        "buys": [
            {
                "buyer": {
                    "id": "0x352c6b116bd19928b487a627b936753f71a7cd1e"
                },
                "sale": {
                    "id": "0x463b5fe6dd5633e6fecb286fdbab9d3a9a75b038/0x1",
                    "seller": {
                        "id": "0xb9602f2442da97651d5f7e0435a4733b1a1145cd"
                    },
                    "nft": {
                        "id": "0xb90dc9e354001e6260de670edd6abadb890c6ac9/0x1",
                        "contract": {
                            "id": "0xb90dc9e354001e6260de670edd6abadb890c6ac9"
                        }
                    }
                },
                "timestamp": "1653938286"
            }, ...
        ]
    }
}`
