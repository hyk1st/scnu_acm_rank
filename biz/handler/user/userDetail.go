package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
)

func UserDetail(ctx context.Context, c *app.RequestContext) {
	userInter, flag := c.Get("user")
	if !flag {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "没有登录",
		})
		return
	}
	user, ok := userInter.(model.User)
	if !ok {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "没有登录",
		})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    user,
	})
}
