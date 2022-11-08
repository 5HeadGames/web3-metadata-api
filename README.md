# Enders API
Go REST API to serve metadata indexed from a subgraph and a database containing on-chain events and user generated data.

## Current Endpoints
### "/"
**Method**: GET
**Returns**: { string: string }
**Response:** `{"message": "Enders API"}`

### "/packs"
**Method**: GET
**Returns**: { data: [erc1155Tokens: []struct {
            totalSupply: struct {
                value: string
            }
        } ] 
    }
