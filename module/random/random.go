package random

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/log"
)

type RType int

const (
	Avatar RType = iota
	Signature
	BackgroundIMG
)

var (
	avatar        = "https://api.vvhan.com/api/avatar?type=json"
	signature     = "https://api.vvhan.com/api/love?type=json"
	backgroundIMG = "https://api.vvhan.com/api/bing?size=400x240&type=json&rand=sj"
)

type AvatarResponse struct {
	Success bool   `json:"success"`
	Avatar  string `json:"avatar"`
}

type Data struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type BGResponse struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type SignatureResponse struct {
	Success bool   `json:"success"`
	Ishan   string `json:"ishan"`
}

func GetAvatarAndBGIMGAndSignature() (string, string, string) {
	avatar, err := Random(Avatar)
	if err != nil {
		log.Logger.Error(err)
		avatar = config.DEFAULT.Avatar
	}

	bgIMG, err := Random(BackgroundIMG)
	if err != nil {
		log.Logger.Error(err)
		bgIMG = config.DEFAULT.BackgroundIMG
	}

	signature, err := Random(Signature)
	if err != nil {
		log.Logger.Error(err)
		signature = config.DEFAULT.Signature
	}
	return avatar, bgIMG, signature
}

// Random 用于获取随机头像，个性签名，背景图片， 参数：类型（ Avatar、Signature、BackgroundIMG）
func Random(t RType) (string, error) {
	switch t {
	case Avatar:
		i, err := req(avatar, t)
		if err != nil {
			return "", err
		}
		return getRes(i)
	case BackgroundIMG:
		i, err := req(backgroundIMG, t)
		if err != nil {
			return "", err
		}
		return getRes(i)
	case Signature:
		i, err := req(signature, t)
		if err != nil {
			return "", err
		}
		return getRes(i)
	}

	return "", errors.New("unknown error")
}

func req(api string, t RType) (interface{}, error) {
	log.Logger.Infof("api: %v", api)

	var resp *http.Response
	resp, err := http.Get(api)
	if err != nil {
		log.Logger.Error(err)
	}

	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error(err)
	}

	log.Logger.Infof("all: %v", string(all))

	switch t {
	case Avatar:
		var respData AvatarResponse
		err = json.Unmarshal(all, &respData)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
		log.Logger.Infof("resp data: %v", respData)
		return respData, nil
	case BackgroundIMG:
		var respData BGResponse
		err = json.Unmarshal(all, &respData)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
		log.Logger.Infof("resp data: %v", respData)
		return respData, nil
	case Signature:
		var respData SignatureResponse
		err = json.Unmarshal(all, &respData)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
		log.Logger.Infof("resp data: %v", respData)
		return respData, nil
	}
	return nil, errors.New("unknown error")
}

func getRes(resp interface{}) (string, error) {
	switch v := resp.(type) {
	case AvatarResponse:
		if v.Success {
			return v.Avatar, nil
		} else {
			return config.DEFAULT.Avatar, nil
		}
	case BGResponse:
		if v.Success {
			return v.Data.URL, nil
		} else {
			return config.DEFAULT.BackgroundIMG, nil
		}
	case SignatureResponse:
		if v.Success {
			return v.Ishan, nil
		} else {
			return config.DEFAULT.Signature, nil
		}
	}
	return "", errors.New("unknown error")
}
