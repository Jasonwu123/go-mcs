package helper

import (
	"fmt"
	"go-mcs/utils_tool"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
)

// Build up http request body by params
func createReqBody(filePath string) (string, io.Reader, error) {
	var err error
	pr, pw := io.Pipe()
	bw := multipart.NewWriter(pw)
	f, err := utils_tool.OpenFile(filePath)
	if err != nil {
		return "", nil, err
	}

	go func() {
		defer f.Close()
		// part1 wallet_address
		p1w, _ := bw.CreateFormField("wallet_address")
		p1w.Write([]byte(WALLET_ADDRESS))

		// part2 duration
		p2w, _ := bw.CreateFormField("duration")
		p2w.Write([]byte(DURATION))

		// part3 file type
		p3w, _ := bw.CreateFormField("file_type")
		p3w.Write([]byte(FILETYPE))

		// part4 file
		_, fileName := filepath.Split(filePath)
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, utils_tool.EscapeQuotes("file"), utils_tool.EscapeQuotes(fileName)))
		h.Set("Content-Type", "application/octet-stream")
		fw1, _ := bw.CreatePart(h)
		cnt, _ := io.Copy(fw1, f)
		log.Printf("copy %d bytes from file %s in total\n", cnt, fileName)
		bw.Close()
		pw.Close()
	}()

	return bw.FormDataContentType(), pr, nil
}

// Upload file for mcs swam server
func UploadFile(filePath string) error {
	// create body
	contType, reader, err := createReqBody(filePath)
	if err != nil {
		return err
	}

	log.Printf("createReqBody ok\n")

	url := fmt.Sprintf("%s/storage/ipfs/upload", MCS_API)

	req, err := http.NewRequest("POST", url, reader)

	// add headers
	req.Header.Add("Content-Type", contType)

	if err != nil {
		fmt.Println("request new error: ", err)
		return err
	}

	client := &http.Client{}
	log.Printf("upload %s...\n", filePath)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request send error: ", err)
		return err
	}

	defer resp.Body.Close()
	log.Printf("upload %s ok\n", filePath)
	//content, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("resp read body error: ", err)
	//	return err
	//}
	//log.Printf("upload file result: %s\n", content)
	utils_tool.ReadResp(resp.Body)

	return nil
}