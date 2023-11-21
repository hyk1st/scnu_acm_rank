package root

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	json2 "github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gorm.io/gorm"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/remote"
	"scnu_acm_rank/biz/reqModel"
)

func UpdateContestResult(ctx context.Context, c *app.RequestContext) {
	req := reqModel.UpdateContestResultReq{}
	err := c.BindForm(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	f := remote.FileCrawler{}
	analysisRes, err := f.AnalysisRes(req.Json)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	comp := model.Competiton{}
	model.DB.Model(&comp).Where("id = ?", comp.Id).Find(&comp)
	err = model.DB.Transaction(func(tx *gorm.DB) error {
		tx.Save(comp)
		for _, v := range analysisRes.Result {
			js, _ := json2.Marshal(v)
			ins := model.UserCompetition{
				VjName:  v.Name,
				CompId:  comp.Id,
				Rank:    uint(v.Rank),
				Penalty: v.Penalty,
				Solve:   v.SolveCnt,
				Ext:     string(js),
			}
			tx.Create(&ins)
		}
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
	})
}
