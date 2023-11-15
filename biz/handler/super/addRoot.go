package super

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
)

func AddRoot(ctx context.Context, c *app.RequestContext) {
	req := reqModel.AddRootReq{}
	err := c.BindForm(&req)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	model.DB.Model(&model.User{}).Where("stu_id = ?", req.StuID).Update("level", req.Level)
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
	})
}
