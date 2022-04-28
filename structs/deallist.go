package structs

type DealList struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Data   []struct {
		Id          int    `json:"id"`
		FileName    string `json:"file_name"`
		FileSize    string `json:"file_size"`
		CreateAt    string `json:"create_at"`
		MinerFid    string `json:"miner_fid"`
		DealStatus  string `json:"deal_status"`
		Status      string `json:"status"`
		PinStatus   string `json:"pin_status"`
		PayloadCid  string `json:"payload_cid"`
		DealCid     string `json:"deal_cid"`
		DealId      int    `json:"deal_id"`
		PieceCid    string `json:"piece_cid"`
		Duration    int    `json:"duration"`
		LockedFee   string `json:"locked_fee"`
		NftTxHash   string `json:"nft_tx_hash"`
		TokenId     string `json:"token_id"`
		MintAddress string `json:"mint_address"`
	} `json:"data"`
	PageInfo struct {
		PageNumber       string `json:"page_number"`
		PageSize         string `json:"page_size"`
		TotalRecordCount string `json:"total_record_count"`
	} `json:"page_info"`
}
