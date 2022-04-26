package helper

import (
	"fmt"
	"go-mcs/utils_tool"
	"log"
	"net/http"
)

func GetParams() map[string]interface{} {
	url := MCS_API + "/common/system/params"
	content, err := utils_tool.HttpGet(url)
	if err != nil {
		log.Println("Get params error: ", err)
		return nil
	}
	return content
}

func GetFileStatus(cid string) map[string]interface{} {
	url := fmt.Sprintf("%s/storage/deal/log/%s", MCS_API, cid)
	content, err := utils_tool.HttpGet(url)
	if err != nil {
		log.Println("Get file status error: ", err)
		return nil
	}
	return content
}

func GetDealDetail(cid string, dealId int64) map[string]interface{} {
	url := fmt.Sprintf("%s/storage/deal/detail/%d?payload_cid=%s", MCS_API, dealId, cid)
	fmt.Println(url)
	content, err := utils_tool.HttpGet(url)
	if err != nil {
		log.Println("Get deal detail error: ", err)
		return nil
	}
	return content
}

func GetPaymentInfo(cid string) map[string]interface{} {
	url := fmt.Sprintf("%s/billing/deal/lockpayment/info?payload_cid=%s", MCS_API, cid)
	content, err := utils_tool.HttpGet(url)
	if err != nil {
		log.Println("Get payment info error: ", err)
		return nil
	}

	return content
}

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

func GetDealList(cid, fileName, pageNumber, pageSize string) map[string]interface{} {
	if pageNumber == "" {
		pageNumber = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}
	if fileName != "" {
		fileName = fileName
	}
	url := fmt.Sprintf("%s/storage/tasks/deals?page_size=%s&page_number=%s&file_name=%s&source_id=4&wallet_address=%s&payload_cid=%s",
		MCS_API, pageSize, pageNumber, fileName, WALLET_ADDRESS, cid,
	)
	content, err := utils_tool.HttpGet(url)
	if err != nil {
		log.Println("Get deal list error: ", err)
		return nil
	}
	return content
}
