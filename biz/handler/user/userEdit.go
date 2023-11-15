package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
)

func EditUser(ctx context.Context, c *app.RequestContext) {
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
	v := reqModel.UserEditReq{}
	err := c.BindForm(&v)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	usr := model.User{}
	model.DB.Model(&model.User{}).Where("stu_id = ?", v.StuId).Find(&usr)
	v.Change2UserModel(&usr)
	model.DB.Save(&usr)
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    user,
	})
}
