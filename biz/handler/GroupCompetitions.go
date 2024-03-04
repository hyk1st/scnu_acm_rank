package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"scnu_acm_rank/biz/middle"
	"scnu_acm_rank/biz/model"
)

func GroupCompetitions(ctx context.Context, c *app.RequestContext) {
	//c.
	v := c.Query("stu_id")
	res, err := model.GetGroupCompetitionsByTeam(v)
	if err != nil {
		c.JSON(http.StatusOK, middle.FailResp(err))
	}
	//model.DB.Raw(sql, v).Find(&temp)
	mp := make([]map[string]interface{}, 0)
	for _, v := range res {
		t := make(map[string]interface{}, 0)
		t["name"] = v.Name
		t["comp_name"] = v.CompName
		t["rank"] = v.Rank
		t["penalty"] = v.Penalty / 1000 / 60
		t["solve"] = v.Solve
		mp = append(mp, t)
	}
	c.JSON(http.StatusOK, middle.SuccessResp("", mp))

}
