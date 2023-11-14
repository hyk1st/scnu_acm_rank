package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
)

func Register(ctx context.Context, c *app.RequestContext) {
	user := reqModel.RegisterReq{}
	err := c.BindForm(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, err)
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
			"result": "false",
			"error":  errStr,
		})
	}
	// 模型转换
	userModel := user.ToDbModle()
	model.DB.Create(userModel)
	c.JSON(200, utils.H{
		"result": "success",
	})

}
