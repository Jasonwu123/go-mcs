package utils_tool

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
	"io"
	"log"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ExitIfErr error handle
func ExitIfErr(err error) {
	if err != nil {
		panic(err)
	}
	return
}

// OpenFile read file
func OpenFile(filePath string) (*os.File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func EscapeQuotes(s string) string {
	quoteEscaper := strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
	return quoteEscaper.Replace(s)
}

// ReadResp read http response body
func ReadResp(resp io.Reader) map[string]interface{} {
	content, err := io.ReadAll(resp)
	if err != nil {
		fmt.Println("resp read body error: ", err)
		return nil
	}
	return byteConvToMap(content)
}

// Byte convert to map
func byteConvToMap(content []byte) map[string]interface{} {
	tmpCon := map[string]interface{}{}
	if err := json.Unmarshal(content, &tmpCon); err != nil {
		log.Println("byte converto to map error: ", err)
		return nil
	}
	return tmpCon
}

// HttpGet info
func HttpGet(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return ReadResp(resp.Body), nil
}

// Multipar http request
func MultipartReq(url string, params map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range params {
		_ = writer.WriteField(k, v)
	}
	writer.Close()
	return http.NewRequest("POST", url, body)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// PublicKeyBytesToAddress ...
func PublicKeyBytesToAddress(publicKey []byte) common.Address {
	var buf []byte

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[1:]) // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}

// interface to uint64
func InterfaceToUint64(v interface{}) (uint64, error) {
	switch t := v.(type) {
	case nil:
		return 0, errors.New("invalid v, not found value")
	case int:
		return uint64(t), nil
	case int16:
		return uint64(t), nil
	case int32:
		return uint64(t), nil
	case int64:
		return uint64(t), nil
	case float32:
		return uint64(t), nil
	case float64:
		return uint64(t), nil
	case uint:
		return uint64(t), nil
	case uint16:
		return uint64(t), nil
	case uint32:
		return uint64(t), nil
	case uint64:
		return t, nil
	case json.Number:
		num, err := t.Int64()
		return uint64(num), err
	case string:
		return strconv.ParseUint(t, 10, 64)
	default:
		return 0, errors.New("invalid num")
	}
}

// interface to *big.Int
func InterfaceToBigInt(v interface{}) (*big.Int, error) {
	switch t := v.(type) {
	case nil:
		return big.NewInt(0), errors.New("invalid v, not found value")
	case int:
		return big.NewInt(int64(t)), nil
	case int16:
		return big.NewInt(int64(t)), nil
	case int32:
		return big.NewInt(int64(t)), nil
	case int64:
		return big.NewInt(int64(t)), nil
	case float32:
		return big.NewInt(int64(t)), nil
	case float64:
		return big.NewInt(int64(t)), nil
	case uint16:
		return big.NewInt(int64(t)), nil
	case uint32:
		return big.NewInt(int64(t)), nil
	case uint64:
		return big.NewInt(int64(t)), nil
	case json.Number:
		num, err := t.Int64()
		return big.NewInt(num), err
	case string:
		i, err := strconv.Atoi(t)
		return big.NewInt(int64(i)), err
	case *big.Int:
		return t, nil
	default:
		return big.NewInt(0), errors.New("invalid num")
	}
}
