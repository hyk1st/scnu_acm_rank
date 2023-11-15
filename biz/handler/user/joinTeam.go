package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
)

func JoinTeam(ctx context.Context, c *app.RequestContext) {
	userInter, flag := c.Get("user")
	if !flag {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "没有登录",
		})
		return
	}
	user, ok := userInter.(model.User)
	if !ok {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "没有登录",
		})
		return
	}
	req := reqModel.JoinTeamReq{}
	err := c.BindForm(&req)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	team := model.Team{Id: 0}
	model.DB.Model(&team).Where("key = ?", req.Key).Find(&team)
	if team.Id == 0 {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "key不存在",
		})
		return
	}
	var cnt int64
	// 检验队伍人数
	model.DB.Model(&user).Find("where group_id = ?", team.Id).Count(&cnt)
	if cnt >= 3 {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "队伍人数已满",
		})
		return
	}
	mutex.Lock()
	model.DB.Model(&user).Update("group_id", team.Id)
	mutex.Unlock()
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
	})
}
