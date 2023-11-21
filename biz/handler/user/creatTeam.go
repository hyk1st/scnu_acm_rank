package user

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gorm.io/gorm"
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
	user := v.(*model.User)
	team := req.GetModel()
	team.Leader = user.StuId
	team.Key = fmt.Sprintf("%v", time.Now().UnixNano())
	team.Status = 1
	err = model.DB.Transaction(func(tx *gorm.DB) error {
		tx.Save(&team).Find(&team)
		tx.Model(&user).Where("stu_id = ?", user.StuId).Update("group_id", team.Id)
		return nil
	})
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    team,
	})
}
