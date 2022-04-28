package helper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mcs/structs"
)

func Mint(client *ethclient.Client, payer common.Address, cid string, nft structs.NFT) {
	//paymentInfo := GetPaymentInfo(cid)
	//txHash := paymentInfo.Data.TxHash
	//
	//nft.TxHash = txHash
	//
	//uploadResonse := UploadFile("", )
}
