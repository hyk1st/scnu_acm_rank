package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/remote"
)

func CompetitionDetail(ctx context.Context, c *app.RequestContext) {
	if _, ok := c.Get("id"); !ok {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "参数错误",
		})
	}
	id, _ := c.Get("id")
	comp := model.Competiton{}
	model.DB.Model(&model.Competiton{}).Where("id = ?", id).First(&comp)
	if len(comp.Result) < 1 {
		var crawler remote.CrawlTrainRes
		crawler.GetTrainRes()
		crawler.AnalysisRes()
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    comp,
	})
}
