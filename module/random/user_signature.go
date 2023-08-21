package random

import (
	"biu-x.org/TikTok/module/log"
	"encoding/json"
	"io"
	"net/http"
)

type SignatureResponse struct {
	Success bool   `json:"success"`
	Ishan   string `json:"ishan"`
}

func GetRandomSignature() (string, error) {
	var respData SignatureResponse
	resp, err := http.Get("https://api.vvhan.com/api/love?type=json")
	if err != nil {
		log.Logger.Error(err)
		// 失败重试
		_, _ = GetRandomSignature()
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Logger.Error(err)
		}
	}(resp.Body)

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error(err)
		return "", err
	}

	log.Logger.Infof("all: %v", string(all))

	err = json.Unmarshal(all, &respData)
	if err != nil {
		log.Logger.Error(err)
		return "", err
	}
	log.Logger.Infof("resp data: %v", respData)

	if respData.Success {
		return respData.Ishan, nil
	}
	return "", nil
}
