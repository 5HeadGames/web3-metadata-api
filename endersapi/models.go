package endersapi

type erc1155Tokens struct {
	Erc1155Tokens []struct {
		TotalSupply struct {
			Value float32 `json:"value"`
		} `json:"totalSupply"`
	} `json:"erc1155Tokens"`
}

type sales struct {
	Sales []struct {
		Seller struct {
			Id string `json:"id"`
		} `json:"seller"`
		Nft struct {
			Contract struct {
				Id string `json:"id"`
			} `json:"contract"`
		} `json:"nft"`
		Timestamp string `json:"timestamp"`
	} `json:"sales"`
}

type buys struct {
	Buys []struct {
		Buyer struct {
			Id string `json:"id"`
		} `json:"buyer"`
		Sale struct {
			Seller struct {
				Id string `json:"id"`
			} `json:"seller"`
		} `json:"sale"`
		Nft struct {
			Contract struct {
				Id string `json:"id"`
			} `json:"contract"`
		} `json:"nft"`
	} `json:"buys"`
}

type PacksResponse struct {
	Data struct{ erc1155Tokens } `json:"data"`
}

type SalesResponse struct {
	Data struct{ sales } `json:"data"`
}

type BuysResponse struct {
	Data struct{ buys } `json:"data"`
}
