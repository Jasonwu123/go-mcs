package helper

import (
	"encoding/json"
	"fmt"
	"go-mcs/structs"
	"go-mcs/utils_tool"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
)

// Create upload file http request body by params
func createReqBody(address, filePath string, fileParams structs.FileParams) (string, io.Reader, error) {
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
		if address == "" {
			address = WALLET_ADDRESS
			p1w, _ := bw.CreateFormField("wallet_address")
			p1w.Write([]byte(address))
		} else {
			p1w, _ := bw.CreateFormField("wallet_address")
			p1w.Write([]byte(address))
		}

		// part2 duration
		if fileParams.Duration == "" {
			p2w, _ := bw.CreateFormField("duration")
			p2w.Write([]byte(DURATION))
		} else {
			p2w, _ := bw.CreateFormField("duration")
			p2w.Write([]byte(fileParams.Duration))
		}

		// part3 file type
		if fileParams.FileType == "" {
			p3w, _ := bw.CreateFormField("file_type")
			p3w.Write([]byte(FILETYPE))
		} else {
			p3w, _ := bw.CreateFormField("file_type")
			p3w.Write([]byte(fileParams.FileType))
		}


		// part4 delay
		if fileParams.Delay == "" {
			p4w, _ := bw.CreateFormField("delay")
			p4w.Write([]byte(DELAY))

		} else {
			p4w, _ := bw.CreateFormField("delay")
			p4w.Write([]byte(fileParams.Delay))
		}

		// part5 file
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

// UploadFile upload file to mcs swam server
func UploadFile(address, filePath string, fileParams structs.FileParams) structs.UploadFileResponse {
	// create body
	contType, reader, err := createReqBody(address, filePath, fileParams)
	if err != nil {
		log.Println(err)
		return structs.UploadFileResponse{}
	}

	log.Printf("createReqBody ok\n")

	url := fmt.Sprintf("%s/storage/ipfs/upload", MCS_API)

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println(err)
		return structs.UploadFileResponse{}
	}

	// add headers
	req.Header.Add("Content-Type", contType)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")

	client := &http.Client{}
	log.Printf("upload %s...\n", filePath)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Upload file failed: ", err)
		return structs.UploadFileResponse{}
	}

	log.Printf("upload %s ok\n", filePath)
	log.Println("")

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return structs.UploadFileResponse{}
	}
	defer resp.Body.Close()

	var responseObject structs.UploadFileResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Println(err)
		return structs.UploadFileResponse{}
	}

	return responseObject
}