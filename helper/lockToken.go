package helper

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mcs/erc20"
	"go-mcs/swanpayment"
	"go-mcs/utils_tool"
	"math/big"
)

func LockToken(client *ethclient.Client, payer common.Address, cid string, amount *big.Int) (*types.Transaction, error) {

	privateKey, err := crypto.HexToECDSA("")
	utils_tool.ExitIfErr(err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	payer = utils_tool.PublicKeyBytesToAddress(publicKeyBytes)

	minAmount := utils_tool.ToWei(0.01, 18)
	amount = utils_tool.ToWei(amount, 18)

	value := big.NewInt(0)

	chainID, err := client.NetworkID(context.Background())
	utils_tool.ExitIfErr(err)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	utils_tool.ExitIfErr(err)

	opts := &bind.TransactOpts{
		From:     payer,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    value,
		GasPrice: auth.GasPrice,
		GasLimit: auth.GasLimit,
		Context:  auth.Context,
		NoSend:   false,
	}

	usdcInstance, err := erc20.NewErc20(common.HexToAddress(USDCAddress), client)
	approveTx, err := usdcInstance.Approve(opts, common.HexToAddress(SWANPaymentAddress), amount)
	utils_tool.ExitIfErr(err)
	_ = approveTx

	paymentInstance, err := swanpayment.NewSwanpayment(common.HexToAddress(SWANPaymentAddress), client)
	utils_tool.ExitIfErr(err)

	param := swanpayment.IPaymentMinimallockPaymentParam{
		Id:         cid,
		MinPayment: minAmount,
		Amount:     amount,
		LockTime:   big.NewInt(int64(86400)).Mul(big.NewInt(int64(86400)), big.NewInt(int64(6))),
		Recipient:  common.HexToAddress(RECIPIENTAddress),
		Size:       big.NewInt(int64(0)),
		CopyLimit:  uint8(1),
	}

	return paymentInstance.LockTokenPayment(opts, param)

}
