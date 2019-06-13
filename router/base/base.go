package base

import (
	"encoding/json"
	"net/http"

	"github.com/yitume/shop-api/model/trans"
	"github.com/yitume/shop-api/pkg/common/code"

	"github.com/gin-gonic/gin"
)

const (
	// MsgOk ajax输出错误码，成功
	MsgOk = 0
	// MsgOk ajax输出错误码，成功
	MsgRedirect = 302
	// MsgErr 错误
	MsgErr = 1
)

// JSONResult json
type JSONResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type JSONResultRaw struct {
	Code    int             `json:"code"`
	Message string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
}

type JSONNewResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}

// JSON 提供了系统标准JSON输出方法。
func JSON(c *gin.Context, Code int, data ...interface{}) {
	result := new(JSONNewResult)
	result.Code = Code
	info, ok := code.CodeMap[Code]
	if ok {
		result.Message = info
	} else {
		result.Message = "error"
	}

	if len(data) > 0 {
		result.Result = data[0]
	} else {
		result.Result = ""
	}
	c.JSON(http.StatusOK, result)
	return
}

// JSON 提供了系统标准JSON输出方法。
func JSONRaw(c *gin.Context, Code int, message string, data []byte) {
	result := new(JSONResultRaw)
	result.Code = Code
	result.Message = message
	result.Data = data
	c.JSON(http.StatusOK, result)
	return
}

// JSON 提供了系统标准JSON输出方法。
func JSONList(c *gin.Context, data interface{}, total int) {
	result := new(JSONNewResult)
	result.Code = 0
	result.Message = "ok"
	result.Result = trans.RespList{
		List:        data,
		TotalNumber: total,
	}
	c.JSON(http.StatusOK, result)
	return
}
