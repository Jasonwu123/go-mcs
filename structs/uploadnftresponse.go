package structs

type NFTResponse struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Data   struct {
		PayloadCid string `json:"payload_cid"`
		IpfsUrl    string `json:"ipfs_url"`
		NeedPay    int    `json:"need_pay"`
	} `json:"data"`
}
