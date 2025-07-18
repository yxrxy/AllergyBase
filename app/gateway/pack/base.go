package pack

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/yxrxy/AllergyBase/pkg/errno"
)

type Base struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type RespWithData struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func RespError(c *app.RequestContext, err error) {
	Errno := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Base{
		Code: strconv.FormatInt(Errno.ErrorCode, 10),
		Msg:  Errno.ErrorMsg,
	})
}

func RespSuccess(c *app.RequestContext) {
	Errno := errno.Success
	c.JSON(consts.StatusOK, Base{
		Code: strconv.FormatInt(Errno.ErrorCode, 10),
		Msg:  Errno.ErrorMsg,
	})
}

func RespData(c *app.RequestContext, items any) {
	Errno := errno.Success
	resp := RespWithData{
		Code: strconv.FormatInt(Errno.ErrorCode, 10),
		Msg:  Errno.ErrorMsg,
		Data: items,
	}
	c.JSON(consts.StatusOK, resp)
}
