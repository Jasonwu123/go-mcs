package helper

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mcs/erc20"
	"go-mcs/swanpayment"
	"go-mcs/utils_tool"
	"math"
	"math/big"
)

/*

	PAY_GAS_LIMIT := params["data"].(map[string]interface{})["PAY_GAS_LIMIT"]
	PAY_WITH_MULTIPLY_FACTOR := params["data"].(map[string]interface{})["PAY_WITH_MULTIPLY_FACTOR"]
	SWAN_PAYMENT_CONTRACT_ADDRESS := params["data"].(map[string]interface{})["SWAN_PAYMENT_CONTRACT_ADDRESS"]
	USDC_ADDRESS := params["data"].(map[string]interface{})["USDC_ADDRESS"]
*/

//func LockToken(client *ethclient.Client, payer common.Address, cid string, amount *big.Int) (*types.Transaction, error) {
//	params := GetParams()
//
//	//usdcAddress := common.HexToAddress(fmt.Sprint(params["data"].(map[string]interface{})["USDC_ADDRESS"]))
//
//	gatewayContractAddress := common.HexToAddress(fmt.Sprint(params["data"].(map[string]interface{})["SWAN_PAYMENT_CONTRACT_ADDRESS"]))
//
//	nonce, err := client.PendingNonceAt(context.TODO(), payer)
//	utils_tool.ExitIfErr(err)
//
//	gasLimitInter := params["data"].(map[string]interface{})["PAY_GAS_LIMIT"]
//
//	gasLimit, err := utils_tool.InterfaceToUint64(gasLimitInter)
//	if err != nil {
//		panic("gasLimit is not the type of uint64.")
//	}
//
//	gasPrice, err := client.SuggestGasPrice(context.TODO())
//	utils_tool.ExitIfErr(err)
//
//	paymentInstance, err := swanpayment.NewSwanpayment(gatewayContractAddress, client)
//
//	opts := &bind.TransactOpts{
//		From:     payer,
//		Nonce:    new(big.Int).SetUint64(nonce),
//		GasPrice: gasPrice,
//		GasLimit: gasLimit,
//	}
//
//	minPayment := utils_tool.ToWei(amount, 18)
//
//	multiplyFactor := params["data"].(map[string]interface{})["PAY_WITH_MULTIPLY_FACTOR"]
//	factor, err := utils_tool.InterfaceToBigInt(multiplyFactor)
//	if err != nil {
//		panic("mutiply factor is not the type of *big.Int")
//	}
//
//	amount.Mul(amount, factor)
//	amount = utils_tool.ToWei(amount, 18)
//
//	lockTimeInter := params["data"].(map[string]interface{})["LOCK_TIME"]
//	lockTime, err := utils_tool.InterfaceToBigInt(lockTimeInter)
//	if err != nil {
//		panic("lockTime is not the type of *big.Int")
//	}
//	lockTime.Mul(lockTime, big.NewInt(86400))
//
//	recipientAddress := common.HexToAddress(fmt.Sprint(params["data"].(map[string]interface{})["RECIPIENT"]))
//
//	param := swanpayment.IPaymentMinimallockPaymentParam{
//		Id:         cid,
//		MinPayment: minPayment,
//		Amount:     amount,
//		LockTime:   lockTime,
//		Recipient:  recipientAddress,
//		Size:       big.NewInt(0),
//		CopyLimit:  1,
//	}
//
//	return paymentInstance.LockTokenPayment(opts, param)
//
//}

func LockToken(client *ethclient.Client, payer common.Address, cid string, amount *big.Int) {
	params := GetParams()

	gatewayContractAddress := common.HexToAddress(fmt.Sprint(params["data"].(map[string]interface{})["SWAN_PAYMENT_CONTRACT_ADDRESS"]))

	nonce, err := client.PendingNonceAt(context.TODO(), payer)
	utils_tool.ExitIfErr(err)

	gasPrice, err := client.SuggestGasPrice(context.TODO())
	utils_tool.ExitIfErr(err)

	paymentInstance, err := swanpayment.NewSwanpayment(gatewayContractAddress, client)

	opts := &bind.TransactOpts{
		From:     payer,
		Nonce:    new(big.Int).SetUint64(nonce),
		GasPrice: gasPrice,
		GasLimit: uint64(9999999),
	}

	minPayment := utils_tool.ToWei(amount, 18)

	multiplyFactor := big.NewInt(int64(math.Floor(1.5)))

	amount.Mul(amount, multiplyFactor)

	lockTime := big.NewInt(6)

	lockTime.Mul(lockTime, big.NewInt(86400))

	recipientAddress := common.HexToAddress(fmt.Sprint(params["data"].(map[string]interface{})["RECIPIENT"]))

	param := swanpayment.IPaymentMinimallockPaymentParam{
		Id:         cid,
		MinPayment: minPayment,
		Amount:     amount,
		LockTime:   lockTime,
		Recipient:  recipientAddress,
		Size:       big.NewInt(0),
		CopyLimit:  1,
	}

	usdcAddress := common.HexToAddress(fmt.Sprint(params["data"].(map[string]interface{})["USDC_ADDRESS"]))
	USDCInstance, err := erc20.NewErc20(usdcAddress, client)
	utils_tool.ExitIfErr(err)

	approveTx, err := USDCInstance.Approve(opts, gatewayContractAddress, amount)
	utils_tool.ExitIfErr(err)

	chainID, err := client.NetworkID(context.TODO())
	utils_tool.ExitIfErr(err)
	fmt.Println("chain id: ", chainID)

	privateKey, err := crypto.HexToECDSA("")
	utils_tool.ExitIfErr(err)

	signedTx, err := types.SignTx(approveTx, types.NewEIP155Signer(chainID), privateKey)
	utils_tool.ExitIfErr(err)

	err = client.SendTransaction(context.TODO(), signedTx)
	utils_tool.ExitIfErr(err)

	fmt.Println("approveTx: ", approveTx.Hash().Hex())
	fmt.Printf("approveTx sent: %s\n", signedTx.Hash().Hex())

	payTx, err := paymentInstance.LockTokenPayment(opts, param)
	utils_tool.ExitIfErr(err)

	signedpayTx, err := types.SignTx(payTx, types.NewEIP155Signer(chainID), privateKey)
	utils_tool.ExitIfErr(err)

	err = client.SendTransaction(context.TODO(), signedpayTx)
	utils_tool.ExitIfErr(err)

}
