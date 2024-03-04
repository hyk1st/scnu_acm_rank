package handler

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
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

	res := make([]map[string]interface{}, 0)
	model.DB.Raw(sql, id).Find(&res)
	if len(res) == 0 {
		c.JSON(http.StatusOK, middle.SuccessResp("", nil))
		return
	}
	list := make([]string, 0, 3)
	for _, v := range res {
		list = append(list, v["userName"].(string))
	}
	c.JSON(http.StatusOK, middle.SuccessResp("", utils.H{
		"name":  res[0]["name"],
		"users": list,
	}))
}
