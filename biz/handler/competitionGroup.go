package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/middle"
	"scnu_acm_rank/biz/model"
	"sort"
)

func CompetitionGroup(ctx context.Context, c *app.RequestContext) {
	res, err := model.GetGroupCompetitions()
	if err != nil {
		c.JSON(http.StatusOK, middle.FailResp(err))
		return
	}
	mp := make(map[int64][]model.GroupResult, 0)
	cnt := make(map[int64]int64, 0)
	for _, row := range res {
		if mp[row.GroupId] == nil || len(mp[row.GroupId]) == 0 {
			mp[row.GroupId] = make([]model.GroupResult, 0)
		}
		if len(mp[row.GroupId]) > 10 {
			continue
		}
		cnt[row.GroupId] += int64(row.Score)
		temp := row
		mp[row.GroupId] = append(mp[row.GroupId], temp)
	}
	slice := make([]int64, 0, len(cnt))
	for k, _ := range cnt {
		slice = append(slice, k)
	}
	sort.Slice(slice, func(a, b int) bool {
		return cnt[slice[a]] > cnt[slice[b]]
	})
	resList := make([]map[string]interface{}, 0, len(slice))
	for i, v := range slice {
		resList = append(resList, map[string]interface{}{
			"rank":   i + 1,
			"name":   mp[v][0].Name,
			"stu_id": mp[v][0].GroupId,
			"score":  cnt[v],
		})
	}
	c.JSON(http.StatusOK, utils.H{
		"status": 0,
		"msg":    "success",
		"data": map[string]interface{}{
			"rank": resList,
		},
	})

}
