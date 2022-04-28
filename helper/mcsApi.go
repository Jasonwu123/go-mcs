package helper

import (
	"encoding/json"
	"fmt"
	"go-mcs/structs"
	"go-mcs/utils_tool"
	"io/ioutil"
	"log"
	"net/http"
)

// GetParams get system config
func GetParams() structs.Params {
	url := MCS_API + "/common/system/params"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get params error: ", err)
		return structs.Params{}
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return structs.Params{}
	}

	var responseObject structs.Params

	err = json.Unmarshal(responseData, &responseObject)

	return responseObject
}

// GetFileStatus get deal logs
func GetFileStatus(cid string) structs.FileStatus {
	url := fmt.Sprintf("%s/storage/deal/log/%s", MCS_API, cid)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get file status error: ", err)
		return structs.FileStatus{}
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return structs.FileStatus{}
	}

	var responseObject structs.FileStatus

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
		return structs.FileStatus{}
	}
	return responseObject
}

// GetDealDetail get deal details
func GetDealDetail(cid string, dealId int64) structs.DealDetail {
	url := fmt.Sprintf("%s/storage/deal/detail/%d?payload_cid=%s", MCS_API, dealId, cid)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get deal detail error: ", err)
		return structs.DealDetail{}
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return structs.DealDetail{}
	}

	var responseObj structs.DealDetail
	json.Unmarshal(responseData, &responseObj)
	return responseObj
}

// GetPaymentInfo get payment info
func GetPaymentInfo(cid string) structs.PaymentInfo {
	url := fmt.Sprintf("%s/billing/deal/lockpayment/info?payload_cid=%s", MCS_API, cid)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get payment info error: ", err)
		return structs.PaymentInfo{}
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return structs.PaymentInfo{}
	}
	defer resp.Body.Close()
	var responseObj structs.PaymentInfo
	err = json.Unmarshal(responseData, &responseObj)
	if err != nil {
		log.Fatal(err)
		return structs.PaymentInfo{}
	}
	return responseObj
}

// PostMintInfo record mint info
func PostMintInfo(mintInfo map[string]string) map[string]interface{} {
	url := MCS_API + "/storage/mint/info"
	request, err := utils_tool.MultipartReq(url, mintInfo)
	if err != nil {
		log.Println("multipart request error: ", err)
		return nil
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		log.Println("post mint info error: ", err)
		return nil
	}
	return utils_tool.ReadResp(resp.Body)
}

// GetDealList get uploaded file list
func GetDealList(cid, fileName, pageNumber, pageSize string) structs.DealList {
	if pageNumber == "" {
		pageNumber = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}
	if fileName != "" {
		fileName = fileName
	} else {
		fileName = ""
	}

	url := fmt.Sprintf("%s/storage/tasks/deals?page_size=%s&page_number=%s&file_name=%s&source_id=4&wallet_address=%s&payload_cid=%s",
		MCS_API, pageSize, pageNumber, fileName, WALLET_ADDRESS, cid,
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get deal list error: ", err)
		return structs.DealList{}
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return structs.DealList{}
	}
	defer resp.Body.Close()

	var responseObj structs.DealList
	err = json.Unmarshal(responseData, &responseObj)
	if err != nil {
		log.Println(err)
		return structs.DealList{}
	}

	return responseObj
}
