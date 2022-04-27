package structs

type Params struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Data   struct {
		LOCKTIME                   string `json:"LOCK_TIME"`
		MINTCONTRACT               string `json:"MINT_CONTRACT"`
		PAYGASLIMIT                string `json:"PAY_GAS_LIMIT"`
		PAYWITHMULTIPLYFACTOR      string `json:"PAY_WITH_MULTIPLY_FACTOR"`
		RECIPIENT                  string `json:"RECIPIENT"`
		SWANPAYMENTCONTRACTADDRESS string `json:"SWAN_PAYMENT_CONTRACT_ADDRESS"`
		USDCADDRESS                string `json:"USDC_ADDRESS"`
	} `json:"data"`
}
