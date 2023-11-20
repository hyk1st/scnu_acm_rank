package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"scnu_acm_rank/biz/model"
	"sort"
)

func CompetitionGroup(ctx context.Context, c *app.RequestContext) {
	res, err := model.GetGroupCompetitions()
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "fail",
			"error":   err,
		})
		return
	}
	mp := make(map[int64][]model.Result, 0)
	cnt := make(map[int64]int64, 0)
	for _, row := range res {
		if mp[row.StuId] == nil || len(mp[row.StuId]) == 0 {
			mp[row.StuId] = make([]model.Result, 0)
		}
		if len(mp[row.StuId]) > 10 {
			continue
		}
		cnt[row.StuId]++
		temp := row
		mp[row.StuId] = append(mp[row.StuId], temp)
	}
	slice := make([]int64, 0, len(cnt))
	for k, _ := range cnt {
		slice = append(slice, k)
	}
	sort.Slice(slice, func(a, b int) bool {
		return cnt[slice[a]] > cnt[slice[b]]
	})

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data": map[string]interface{}{
			"rank":     slice,
			"scoreSum": cnt,
			"detail":   mp,
		},
	})

}
