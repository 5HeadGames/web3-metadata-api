package endersapi

type erc1155Tokens struct {
	Erc1155Tokens []struct {
		TotalSupply struct {
			Value string `json:"value"`
		} `json:"totalSupply"`
	} `json:"erc1155Tokens"`
}

type sales struct {
	Sales []struct {
		Amount   string `json:"amount"`
		Duration string `json:"duration"`
		Nft      struct {
			Id       string `json:"id"`
			Contract struct {
				Id string `json:"id"`
			} `json:"contract"`
		} `json:"nft"`
		Price  string `json:"price"`
		Seller struct {
			Id string `json:"id"`
		} `json:"seller"`
		Timestamp string `json:"timestamp"`
		Status    string `json:"status"`
	} `json:"sales"`
}

type buys struct {
	Buys []struct {
		Buyer struct {
			Id string `json:"id"`
		} `json:"buyer"`
		Sale struct {
			Id     string `json:"id"`
			Seller struct {
				Id string `json:"id"`
			} `json:"seller"`
			Nft struct {
				Id       string `json:"id"`
				Contract struct {
					Id string `json:"id"`
				} `json:"contract"`
			} `json:"nft"`
		} `json:"sale"`
		Timestamp string `json:"timestamp"`
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
