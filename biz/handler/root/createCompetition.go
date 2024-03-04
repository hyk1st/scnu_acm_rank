package root

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	json2 "github.com/cloudwego/hertz/pkg/common/json"
	"gorm.io/gorm"
	"net/http"
	"scnu_acm_rank/biz/middle"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/remote"
	"scnu_acm_rank/biz/reqModel"
)

func CreateCompetition(ctx context.Context, c *app.RequestContext) {
	req := reqModel.CreateContestReq{}
	err := c.BindForm(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, middle.FailResp(err))
		return
	}

	var crawler remote.CrawlTrainRes
	switch req.Kind {
	case 0:
		crawler = remote.VJCrawler
		break
	case 1:
		crawler = remote.NcCrawler
	}

	resJson, err := crawler.GetTrainRes(req.CpId)
	if err != nil {
		c.JSON(http.StatusOK, middle.FailResp(err))
		return
	}
	res, m, err := crawler.AnalysisRes(resJson)
	if err != nil {
		c.JSON(http.StatusOK, middle.FailResp(err))
		return
	}
	m.CpId = req.CpId
	m.Result = resJson
	m.Name = req.Name
	m.BestSolve = res.Result[0].SolveCnt
	m.BestPenalty = res.Result[0].Penalty
	err = model.DB.Transaction(func(tx *gorm.DB) error {

		tx.Save(&m)
		for _, v := range res.Result {
			js, _ := json2.Marshal(v)
			ins := model.UserCompetition{
				VjName:  v.Name,
				CompId:  m.Id,
				Rank:    uint(v.Rank),
				Penalty: v.Penalty,
				Solve:   v.SolveCnt,
				Score:   int(100 - 100*v.Rank/res.Result[len(res.Result)-1].Rank),
				Ext:     string(js),
			}
			tx.Create(&ins)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, middle.FailResp(err))
		return
	}
	c.JSON(http.StatusOK, middle.SuccessResp("创建成功", nil))
}
