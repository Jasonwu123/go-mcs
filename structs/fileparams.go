package structs

type FileParams struct {
	Delay    int    `json:"delay"`
	Duration int    `json:"duration"`
	FileType string `json:"file_type"`
	Address  string `json:"wallet_address"`
}
