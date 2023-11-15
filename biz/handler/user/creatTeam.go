package user

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"runtime"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
	"sync"
	"time"
)

var mutex sync.Mutex

func init() {
	mutex = sync.Mutex{}
	runtime.KeepAlive(mutex)
}

func CreateTeam(ctx context.Context, c *app.RequestContext) {
	req := reqModel.CreateTeamReq{}
	err := c.BindForm(&req)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	v, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   "没有登录",
		})
		return
	}
	user := v.(model.User)
	team := req.GetModel()
	team.Leader = user.Id
	team.Key = fmt.Sprintf("%v", time.Now().Unix())
	team.Status = 1
	mutex.Lock()
	model.DB.Save(&team)
	model.DB.Where("stu_id = ?", user.StuId).Update("group_id", team.Id)
	mutex.Unlock()
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    team,
	})
}
