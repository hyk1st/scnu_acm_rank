package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	json2 "github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gorm.io/gorm"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/remote"
	"scnu_acm_rank/biz/reqModel"
	"scnu_acm_rank/biz/respModel"
)

func CompetitionDetail(ctx context.Context, c *app.RequestContext) {
	req := reqModel.ContestDetailReq{}
	err := c.BindForm(&req)
	if err != nil {
		return
	}
	comp := model.Competition{}
	analysisRes := &remote.AnalysisRes{}
	model.DB.Model(&model.Competition{}).Where("id = ?", req.ContestId).First(&comp)
	if true {
		var crawler remote.CrawlTrainRes = remote.VJCrawler
		res, str, err := crawler.GetTrainRes(comp.VjCpId)
		if err != nil {
			c.JSON(http.StatusOK, utils.H{
				"message": "fail",
				"error":   err,
			})
			return
		}
		analysisRes, err = crawler.AnalysisRes(res)
		if err != nil {
			c.JSON(http.StatusOK, utils.H{
				"message": "fail",
				"error":   err,
			})
			return
		}
		json, err := json2.Marshal(analysisRes)
		if err != nil {
			c.JSON(http.StatusOK, utils.H{
				"message": "fail",
				"error":   err,
			})
			return
		}
		ext := map[string]string{
			"originResult": str,
		}
		extJson, _ := json2.Marshal(ext)
		comp.Ext = string(extJson)
		comp.Result = string(json)
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
	}
	resp := respModel.ContestDetailResp{
		Name:      comp.Name,
		VjCpId:    comp.VjCpId,
		Kind:      comp.Kind,
		Password:  comp.Password,
		StartDate: comp.StartDate,
		Length:    comp.Length,
		CreateUsr: comp.CreateUsr,
		Result:    *analysisRes,
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    resp,
	})
}
