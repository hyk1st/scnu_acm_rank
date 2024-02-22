package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/middle"
	"scnu_acm_rank/biz/reqModel"
)

func SendEmail(ctx context.Context, c *app.RequestContext) {
	req := reqModel.SendEmailReq{}
	err := c.BindForm(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.H{
			"msg":    "fail",
			"status": 1,
			"data":   "",
		})
		return
	}
	err = middle.SendEmail([]string{req.Email})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.H{
			"msg":    "fail",
			"status": 1,
			"data":   "",
		})
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"msg":    "success",
		"status": 0,
		"data":   "",
	})
}
