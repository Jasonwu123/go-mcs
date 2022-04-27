package structs

type PaymentInfo struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Data   struct {
		LockedFee  string `json:"locked_fee"`
		PayloadCid string `json:"payload_cid"`
		TxHash     string `json:"tx_hash"`
	} `json:"data"`
}
