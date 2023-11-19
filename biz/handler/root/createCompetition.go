package root

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
)

func CreateCompetition(ctx context.Context, c *app.RequestContext) {
	req := reqModel.CreateContestReq{}
	err := c.BindForm(&req)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	m := req.Convert2model()
	model.DB.Save(&m)
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
	})
}
