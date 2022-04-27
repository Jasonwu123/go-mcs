package main

import (
	"go-mcs/helper"
)

func main() {
	//rpcUrl := "https://polygon-mumbai.g.alchemy.com/v2/SxI1sWpD8WGsHMcnX3gJajautgx4-vjc"
	//client, err := ethclient.Dial(rpcUrl)
	//utils_tool.ExitIfErr(err)
	//payer := common.HexToAddress(helper.WALLET_ADDRESS)
	//cid := "bafykbzacedesi2mu3fgfrxyh5jkgyiqcmgq2ekz4lblwd7mr5jmu2yyikmznu"
	//amount := big.NewInt(int64(math.Floor(0.05)))
	//helper.LockToken(client, payer, cid, amount)
	cid := "bafykbzacedesi2mu3fgfrxyh5jkgyiqcmgq2ekz4lblwd7mr5jmu2yyikmznu"
	paymentInfo := helper.GetPaymentInfo(cid)

}
