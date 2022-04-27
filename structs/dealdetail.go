package structs

type DealDetail struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Data   struct {
		Dao []struct {
			DaoName         string `json:"dao_name"`
			DaoAddress      string `json:"dao_address"`
			OrderIndex      string `json:"order_index"`
			DealId          int    `json:"deal_id"`
			DaoPassTime     string `json:"dao_pass_time"`
			PayloadCid      string `json:"payload_cid"`
			DaoAddressEvent string `json:"dao_address_event"`
			TxHash          string `json:"tx_hash"`
			Status          string `json:"status"`
		} `json:"dao"`
		DaoThreshHold int `json:"dao_thresh_hold"`
		DaoTotalCount int `json:"dao_total_count"`
		Deal          struct {
			DealId                   int    `json:"deal_id"`
			DealCid                  string `json:"deal_cid"`
			MessageCid               string `json:"message_cid"`
			Height                   int    `json:"height"`
			PieceCid                 string `json:"piece_cid"`
			VerifiedDeal             bool   `json:"verified_deal"`
			StoragePricePerEpoch     int    `json:"storage_price_per_epoch"`
			Signature                string `json:"signature"`
			SignatureType            string `json:"signature_type"`
			CreatedAt                int    `json:"created_at"`
			PieceSizeFormat          string `json:"piece_size_format"`
			StartHeight              int    `json:"start_height"`
			EndHeight                int    `json:"end_height"`
			Client                   string `json:"client"`
			ClientCollateralFormat   string `json:"client_collateral_format"`
			Provider                 string `json:"provider"`
			ProviderTag              string `json:"provider_tag"`
			VerifiedProvider         int    `json:"verified_provider"`
			ProviderCollateralFormat string `json:"provider_collateral_format"`
			Status                   int    `json:"status"`
			NetworkName              string `json:"network_name"`
			StoragePrice             int    `json:"storage_price"`
			IpfsUrl                  string `json:"ipfs_url"`
			FileName                 string `json:"file_name"`
		} `json:"deal"`
		Found struct {
			PayloadCid          string `json:"payload_cid"`
			ClientWalletAddress string `json:"client_wallet_address"`
			CreateAt            string `json:"create_at"`
			LockedFee           string `json:"locked_fee"`
		} `json:"found"`
		SignedDaoCount int  `json:"signed_dao_count"`
		UnlockStatus   bool `json:"unlock_status"`
	} `json:"data"`
}
