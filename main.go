package main

import (
	"fmt"
	"go-mcs/helper"
)

func main() {

	//filePath := "/home/jason/Documents/Django.pdf"
	//err := helper.UploadFile(filePath)
	//utils_tool.ExitIfErr(err)

	params := helper.GetParams()
	//fmt.Println(params)

	LOCK_TIME := params["data"].(map[string]interface{})["LOCK_TIME"]
	fmt.Println("lock time: ", LOCK_TIME)

	MINT_CONTRACT := params["data"].(map[string]interface{})["MINT_CONTRACT"]
	fmt.Println("mint contract: ", MINT_CONTRACT)

	PAY_GAS_LIMIT := params["data"].(map[string]interface{})["PAY_GAS_LIMIT"]
	PAY_WITH_MULTIPLY_FACTOR := params["data"].(map[string]interface{})["PAY_WITH_MULTIPLY_FACTOR"]
	SWAN_PAYMENT_CONTRACT_ADDRESS := params["data"].(map[string]interface{})["SWAN_PAYMENT_CONTRACT_ADDRESS"]
	USDC_ADDRESS := params["data"].(map[string]interface{})["USDC_ADDRESS"]

	fmt.Println("pay gas limit: ", PAY_GAS_LIMIT)
	fmt.Println("pay with multiply factor: ", PAY_WITH_MULTIPLY_FACTOR)
	fmt.Println("swan payment contract address: ", SWAN_PAYMENT_CONTRACT_ADDRESS)
	fmt.Println("usdc address: ", USDC_ADDRESS)

	//cid := "bafykbzacebaxluhozonmakj7f2qidkufg7mrdhsinw72hinqp7zk5kw2gax5c"
	//content := helper.GetFileStatus(cid)
	//fmt.Println("file status: ", content)

	//cid := "bafykbzacebaxluhozonmakj7f2qidkufg7mrdhsinw72hinqp7zk5kw2gax5c"
	//content := helper.GetDealList(cid, "", "", "")
	//deadId := content["data"].([]interface{})[0].(map[string]interface{})["deal_id"].(float64)
	//deadDetail := helper.GetDealDetail(cid, int64(deadId))
	//fmt.Printf("dead detail: %s", deadDetail)

}
