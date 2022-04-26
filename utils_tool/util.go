package utils_tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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

//
