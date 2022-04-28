package structs

type FileParams struct {
	Delay    string    `json:"delay"`
	Duration string    `json:"duration"`
	FileType string `json:"file_type"`
	Address  string `json:"wallet_address"`
}
