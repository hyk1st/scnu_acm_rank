package root

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/config"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
)

func UpdateConfig(ctx context.Context, c *app.RequestContext) {
	conf := reqModel.UpdateConfigReq{}
	err := c.BindForm(&conf)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	m := conf.Convert2DbModel()
	model.DB.Model(m).Where("1 = ?", 1).Updates(m)
	config.Update <- struct{}{}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
	})
}
