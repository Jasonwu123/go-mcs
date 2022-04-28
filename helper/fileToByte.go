package helper

import (
	"bytes"
	"fmt"
	"go-mcs/utils_tool"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

// Create upload file http request body by params
func createReqBodyTest(filePath string) (string, io.Reader, error) {
	//var err error
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)

	// part1 wallet_address
	p1w, _ := bw.CreateFormField("wallet_address")
	p1w.Write([]byte(WALLET_ADDRESS))

	// part2 duration
	p2w, _ := bw.CreateFormField("duration")
	p2w.Write([]byte(DURATION))

	// part3 file type
	p3w, _ := bw.CreateFormField("file_type")
	p3w.Write([]byte(NFTFILETYPE))

	// part4 file
	//_, fileName := filepath.Split(filePath)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, utils_tool.EscapeQuotes("file"), utils_tool.EscapeQuotes("nftName")))
	h.Set("Content-Type", "application/octet-stream")
	fw1, _ := bw.CreatePart(h)
	fw1.Write([]byte(filePath))
	
	defer bw.Close()
	return bw.FormDataContentType(), buf, nil
}

// UploadFile upload file to mcs swam server
func UploadFileTest(filePath string)  {
	// create body
	contType, reader, err := createReqBodyTest(filePath)
	if err != nil {
		log.Println(err)
		return
	}

	url := fmt.Sprintf("%s/storage/ipfs/upload", MCS_API)

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println(err)
		return
	}

	// add headers
	req.Header.Add("Content-Type", contType)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("文件上传失败")
		log.Println(err)
		return
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("responseData: ", string(responseData))

	//var responseObject structs.UploadFileResponse
	//err = json.Unmarshal(responseData, &responseObject)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//return responseObject
}
