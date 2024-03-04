package user

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"scnu_acm_rank/biz/middle"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
)

func JoinTeam(ctx context.Context, c *app.RequestContext) {
	userInter, flag := c.Get("user")
	if !flag {
		c.JSON(http.StatusOK, middle.FailResp(errors.New("没有登录")))
		return
	}
	user, ok := userInter.(*model.User)
	if !ok {
		c.JSON(http.StatusOK, middle.FailResp(errors.New("没有登录")))
		return
	}
	req := reqModel.JoinTeamReq{}
	err := c.BindForm(&req)
	if err != nil {
		c.JSON(http.StatusOK, middle.FailResp(err))
		return
	}
	team := model.Team{}
	model.DB.Model(&team).Where("team.key = ? AND team.leader = ?", req.Key, req.StuId).First(&team)
	if team.Id == 0 {
		c.JSON(http.StatusOK, middle.FailResp(errors.New("不存在该队伍")))
		return
	}
	var cnt int64
	// 检验队伍人数
	model.DB.Model(&user).Where("group_id = ?", team.Id).Count(&cnt)
	if cnt >= 3 {
		c.JSON(http.StatusOK, middle.FailResp(errors.New("队伍人数已满")))
		return
	}
	mutex.Lock()
	model.DB.Model(&user).Where("stu_id = ?", user.StuId).Update("group_id", team.Id)
	mutex.Unlock()
	c.JSON(http.StatusOK, middle.SuccessResp("注册成功", map[string]interface{}{
		"group_id": team.Id,
	}))
}
