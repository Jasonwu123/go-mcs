package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-mcs/structs"
	"io/ioutil"
	"log"
	"net/http"
)

func UloadNFT(nft structs.NFT)  {
	postBody, _ := json.Marshal(nft)
	requestBody := bytes.NewBuffer(postBody)
	url := fmt.Sprintf("%s/storage/ipfs/upload", MCS_API)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(url, "application/json", requestBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
