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
)

func CompetitionDetail(ctx context.Context, c *app.RequestContext) {
	if _, ok := c.Get("id"); !ok {
		c.JSON(
			http.StatusOK, utils.H{
				"message": "fail",
				"error":   "参数错误",
			})
	}
	id, _ := c.Get("id")
	comp := model.Competiton{}
	model.DB.Model(&model.Competiton{}).Where("id = ?", id).First(&comp)
	if len(comp.Result) < 1 {
		var crawler remote.CrawlTrainRes = remote.VJCrawler
		res, str, err := crawler.GetTrainRes(id.(string))
		if err != nil {
			c.JSON(http.StatusOK, utils.H{
				"message": "fail",
				"error":   err,
			})
			return
		}
		analysisRes, err := crawler.AnalysisRes(res)
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
				user := model.User{}
				tx.Where("vj_name = ?", v.Name).Find(&user)
				js, _ := json2.Marshal(v)
				ins := model.UserCompetiton{
					UserId:    user.Id,
					UserName:  user.Name,
					CompName:  comp.Name,
					CompId:    comp.Id,
					Goal:      v.Rank,
					StartTime: comp.StartDate,
					Ext:       string(js),
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
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    comp,
	})
}
