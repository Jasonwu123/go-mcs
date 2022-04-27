package structs

type FileStatus struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Data   struct {
		OfflineDealLogs []struct {
			Id       int    `json:"id"`
			DealCid  string `json:"deal_cid"`
			Status   string `json:"status"`
			Message  string `json:"message"`
			CreateAt string `json:"create_at"`
		} `json:"offline_deal_logs"`
	} `json:"data"`
}
