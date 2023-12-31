package handler

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
	user, ok := userInter.(*model.User)
	if !ok {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "没有登录",
		})
		return
	}
	model.DB.Model(&user).Where("stu_id = ?", user.StuId).Find(&user)
	comp := make([]model.UserCompetition, 0)
	model.DB.Model(&model.UserCompetition{}).Where("vj_name = ?", user.VjName).Find(&comp)
	user.Password = ""
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data": utils.H{
			"user":    user,
			"contest": comp,
		},
	})
}
