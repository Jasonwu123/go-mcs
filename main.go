package main

import (
	"fmt"
	"go-mcs/helper"
	"go-mcs/structs"
)

func main() {
	//rpcUrl := "https://polygon-mumbai.g.alchemy.com/v2/SxI1sWpD8WGsHMcnX3gJajautgx4-vjc"
	//client, err := ethclient.Dial(rpcUrl)
	//utils_tool.ExitIfErr(err)
	//payer := common.HexToAddress(helper.WALLET_ADDRESS)
	//cid := "bafykbzacedesi2mu3fgfrxyh5jkgyiqcmgq2ekz4lblwd7mr5jmu2yyikmznu"
	//amount := big.NewInt(int64(math.Floor(0.05)))
	//helper.LockToken(client, payer, cid, amount)
	//cid := "bafykbzacedesi2mu3fgfrxyh5jkgyiqcmgq2ekz4lblwd7mr5jmu2yyikmznu"
	//paymentInfo := helper.GetPaymentInfo(cid)
	//
	//utils_tool.PrintMap(paymentInfo)

	//helper.UploadFile("./tmp.txt")
	//helper.GetParams()
	//filestatus := helper.GetFileStatus(cid)
	//fmt.Println(filestatus.Data.OfflineDealLogs)

	//paymentinfo := helper.GetPaymentInfo(cid)
	//fmt.Println(paymentinfo.Data.PayloadCid)
	//fmt.Println(paymentinfo.Data.TxHash)

	//dealdetail := helper.GetDealDetail(cid, 1)
	//fmt.Println(dealdetail.Data.Deal.FileName)

	//deallist := helper.GetDealList(cid, "", "", "")
	//fmt.Println("code: ", deallist.Code)
	//fmt.Println("status: ", deallist.Status)
	//
	//for i := range deallist.Data {
	//	fmt.Println(deallist.Data[i])
	//}

	fileparams := structs.FileParams{
		Duration: "180",
		FileType: "0",
	}

	filepath := "/Users/jasonwu/Documents/中国史纲.xmind"

	resp := helper.UploadFile("", filepath, fileparams)
	fmt.Println(resp)
	fmt.Println("")
	fmt.Println(resp.Code)
	fmt.Println("")
	fmt.Println(resp.Status)
	fmt.Println("")
	fmt.Println(resp.Data.PayloadCid)
	fmt.Println("")
	fmt.Println(resp.Data.IpfsUrl)

}
