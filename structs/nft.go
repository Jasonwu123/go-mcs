package structs

type NFT struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image      string `json:"image"`
	TxHash string `json:"tx_hash"`
	Duration string `json:"duration"`
	FileType string `json:"file_type"`
	WalletAddress string `json:"wallet_address"`
}
