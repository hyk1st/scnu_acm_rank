package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
)

func TeamDetail(ctx context.Context, c *app.RequestContext) {
	tid, ok := c.Get("teamId")
	if !ok {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "参数错误",
		})
	}
	team := &model.Team{}
	model.DB.Model(team).Where("group_id = ?", tid).Find(&team)
	comp := make([]model.UserCompetiton, 0)
	model.DB.Model(&model.UserCompetiton{}).Where("stu_id = ?", team.Id).Find(&comp)
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data": utils.H{
			"team":    team,
			"contest": comp,
		},
	})
}
