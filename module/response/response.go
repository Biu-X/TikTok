package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func ErrRespWithData(c *gin.Context, msg string, data map[string]interface{}) {
	resp := map[string]interface{}{
		"status_code": Error,
		"status_msg":  msg,
	}
	for key, value := range data {
		resp[key] = value
	}
	c.JSON(http.StatusOK, resp)
}
