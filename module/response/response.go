package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrReponse struct {
	StatusCode int    `json:"status_code"` // 错误码
	Message    string `json:"messgae"`     // 错误信息
}

func OKResp(c *gin.Context) {
	resp := map[string]interface{}{
		"status_code": OK,
		"status_msg":  Msg(OK),
	}
	c.JSON(http.StatusOK, resp)
}

func OKRespWithData(c *gin.Context, data map[string]interface{}) {
	resp := map[string]interface{}{
		"status_code": OK,
		"status_msg":  Msg(OK),
	}
	for key, value := range data {
		resp[key] = value
	}
	c.JSON(http.StatusOK, resp)
}

func ErrResp(c *gin.Context) {
	resp := map[string]interface{}{
		"status_code": Error,
		"status_msg":  Msg(Error),
	}
	c.JSON(http.StatusOK, resp)
}

func ErrRespWithMsg(c *gin.Context, msg string) {
	resp := map[string]interface{}{
		"status_code": Error,
		"status_msg":  msg,
	}
	c.JSON(http.StatusOK, resp)
}
