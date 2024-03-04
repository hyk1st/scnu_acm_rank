package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/middle"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
)

func Register(ctx context.Context, c *app.RequestContext) {
	user := reqModel.RegisterReq{}
	err := c.BindForm(&user)
	if err != nil {
		//fmt.Println(err)
		c.JSON(http.StatusOK, utils.H{
			"status": 1,
			"msg":    "数据格式不正确",
			"data":   "",
		})
		return
	}
	if middle.GetCode(user.Email) != user.Code {
		c.JSON(http.StatusOK, utils.H{
			"status": 1,
			"msg":    "验证码错误",
			"data":   "",
		})
		return
	}
	// 参数校验
	var cnt int64
	errStr := make([]string, 0)
	model.DB.Model(&model.User{}).Where("stu_id = ?", user.StuId).Count(&cnt)
	if cnt > 0 {
		errStr = append(errStr, "用户已存在")
	}

	user.RegisterCheck(&errStr)

	if len(errStr) > 0 {
		c.JSON(http.StatusOK, utils.H{
			"status": 1,
			"msg":    errStr[0],
			"data":   "",
		})
		return
	}
	// 模型转换
	userModel := user.ToDbModle()
	model.DB.Create(userModel)
	c.JSON(200, utils.H{
		"status": 0,
		"msg":    "注册成功",
		"data":   "",
	})

}
