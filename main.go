package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mcs/helper"
	"go-mcs/utils_tool"
	"math/big"
)

func main() {

	//filePath := "/home/jason/Documents/Django.pdf"
	//err := helper.UploadFile(filePath)
	//utils_tool.ExitIfErr(err)

	//params := helper.GetParams()
	////fmt.Println(params)
	//
	//LOCK_TIME := params["data"].(map[string]interface{})["LOCK_TIME"]
	//fmt.Println("lock time: ", LOCK_TIME)
	//
	//MINT_CONTRACT := params["data"].(map[string]interface{})["MINT_CONTRACT"]
	//fmt.Println("mint contract: ", MINT_CONTRACT)
	//
	//PAY_GAS_LIMIT := params["data"].(map[string]interface{})["PAY_GAS_LIMIT"]
	//PAY_WITH_MULTIPLY_FACTOR := params["data"].(map[string]interface{})["PAY_WITH_MULTIPLY_FACTOR"]
	//SWAN_PAYMENT_CONTRACT_ADDRESS := params["data"].(map[string]interface{})["SWAN_PAYMENT_CONTRACT_ADDRESS"]
	//USDC_ADDRESS := params["data"].(map[string]interface{})["USDC_ADDRESS"]
	//
	//fmt.Println("pay gas limit: ", PAY_GAS_LIMIT)
	//fmt.Println("pay with multiply factor: ", PAY_WITH_MULTIPLY_FACTOR)
	//fmt.Println("swan payment contract address: ", SWAN_PAYMENT_CONTRACT_ADDRESS)
	//fmt.Println("usdc address: ", USDC_ADDRESS)

	//cid := "bafykbzacebaxluhozonmakj7f2qidkufg7mrdhsinw72hinqp7zk5kw2gax5c"
	//content := helper.GetFileStatus(cid)
	//fmt.Println("file status: ", content)

	//cid := "bafykbzacebaxluhozonmakj7f2qidkufg7mrdhsinw72hinqp7zk5kw2gax5c"
	//content := helper.GetDealList(cid, "", "", "")
	//deadId := content["data"].([]interface{})[0].(map[string]interface{})["deal_id"].(float64)
	//deadDetail := helper.GetDealDetail(cid, int64(deadId))
	//fmt.Printf("dead detail: %s", deadDetail)

	// mumbai rpc
	rpcUrl := "https://polygon-mumbai.g.alchemy.com/v2/SxI1sWpD8WGsHMcnX3gJajautgx4-vjc"
	client, err := ethclient.Dial(rpcUrl)
	utils_tool.ExitIfErr(err)
	//
	//header, err := client.HeaderByNumber(context.TODO(), nil)
	//utils_tool.ExitIfErr(err)
	//
	//fmt.Println("header number: ", header.Number.String())
	//
	//blockNumber := header.Number
	//
	//block, err := client.BlockByNumber(context.TODO(), blockNumber)
	//utils_tool.ExitIfErr(err)
	//
	//for _, tx := range block.Transactions() {
	//	fmt.Println("tx hash: ", tx.Hash().Hex())
	//	fmt.Println("tx value: ", tx.Value().String())
	//	fmt.Println("tx gas: ", tx.Gas())
	//	fmt.Println("tx gas price: ", tx.GasPrice().Uint64())
	//	fmt.Println("tx nonce: ", tx.Nonce())
	//	//fmt.Printf("tx data: %s\n", string(tx.Data()))
	//	fmt.Println("tx to address: ", tx.To().Hex())
	//
	//	chainID, err := client.NetworkID(context.TODO())
	//	utils_tool.ExitIfErr(err)
	//
	//	if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err == nil {
	//		fmt.Println("tx from address: ", msg.From().Hex())
	//	}
	//	receipt, err := client.TransactionReceipt(context.TODO(), tx.Hash())
	//	utils_tool.ExitIfErr(err)
	//	fmt.Println("receipt status: ", receipt.Status)
	//}
	privateKey, err := crypto.HexToECDSA("")
	utils_tool.ExitIfErr(err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	payer := utils_tool.PublicKeyBytesToAddress(publicKeyBytes)

	cid := "bafykbzacebaxluhozonmakj7f2qidkufg7mrdhsinw72hinqp7zk5kw2gax5c"
	amount := big.NewInt(1)
	helper.LockToken(client, payer, cid, amount)
}
