package helper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Mint(client *ethclient.Client, payer common.Address, cid string, nft struct{}) {
	paymentInfo := GetPaymentInfo(cid)
	_ = paymentInfo
}
