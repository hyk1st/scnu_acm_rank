package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"scnu_acm_rank/biz/middle"
	"scnu_acm_rank/biz/model"
)

func UserCompetitions(ctx context.Context, c *app.RequestContext) {
	//c.
	v := c.Query("stu_id")
	sql := `select b.name, b.stu_id, a.rank, a.solve, ceil(a.penalty / 1000 / 60) penalty, c.name comp_name from user_competition a, user b, competition c where b.stu_id = ? AND  c.kind < 2 AND a.comp_id = c.id AND (a.vj_name = b.vj_name OR a.vj_name = b.nc_name) order BY b.stu_id, c.start_date limit 10`
	temp := make([]map[string]interface{}, 0)
	model.DB.Raw(sql, v).Find(&temp)
	c.JSON(http.StatusOK, middle.SuccessResp("", temp))

}
