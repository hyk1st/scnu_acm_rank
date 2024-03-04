package handler

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"scnu_acm_rank/biz/middle"
	"scnu_acm_rank/biz/model"
)

func TeamDetail(ctx context.Context, c *app.RequestContext) {
	id := c.Query("stu_id")
	tp := c.Query("type")
	if len(id) == 0 || (tp != "person" && tp != "group") {
		c.JSON(http.StatusOK, middle.FailResp(errors.New("类型错误")))
	}
	sql := ""
	if tp == "person" {
		sql = `SELECT team.name, user.name userName from user, team where user.group_id = team.id AND team.id = (SELECT t.group_id from user t where t.stu_id = ?)`
	} else {
		sql = `SELECT user.name userName, team.name from user, team where team.id = ? AND user.group_id = team.id;`
	}
	res := make(map[string]interface{})
	model.DB.Raw(sql, id).Find(res)
	c.JSON(http.StatusOK, middle.SuccessResp("", res))
}
