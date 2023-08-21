package random

import (
	"biu-x.org/TikTok/module/log"
	"encoding/json"
	"io"
	"net/http"
)

type Data struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type BGResponse struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

func GetRandomBackgroundImg() (string, error) {
	var respData BGResponse
	resp, err := http.Get("https://api.vvhan.com/api/bing?size=400x240&type=json&rand=sj")
	if err != nil {
		log.Logger.Error(err)
		// 失败重试
		_, _ = GetRandomBackgroundImg()
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
		return respData.Data.URL, nil
	}
	return "", nil
}
