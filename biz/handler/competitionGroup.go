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
		cnt[row.StuId] += row.Goal
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
		"user":    c.Get("user"),
		"data": map[string]interface{}{
			"rank":     slice,
			"scoreSum": cnt,
			"detail":   mp,
		},
	})
	//groupRes := make(map[int]map[string]model.Result)
	//for _, row := range res {
	//	if _, ok := groupRes[row.GroupId]; !ok {
	//		groupRes[row.GroupId] = make(map[string]model.Result, 0)
	//	}
	//	if groupRes[row.GroupId][row.CompName].Goal < row.Goal {
	//		temp := row
	//		groupRes[row.GroupId][row.CompName] = temp
	//	}
	//}
	//mp := make(map[int][]model.Result, 0)
	//cnt := make(map[int]int64, 0)
	//for k, v := range groupRes {
	//	if mp[k] == nil || len(mp[k]) == 0 {
	//		mp[k] = make([]model.Result, 0)
	//	}
	//	for _, v1 := range v {
	//		temp := v1
	//		mp[k] = append(mp[k], temp)
	//	}
	//	sort.Slice(mp[k], func(a, b int) bool {
	//		return mp[k][a].Goal > mp[k][a].Goal
	//	})
	//	if len(mp[k]) > 10 {
	//		mp[k] = mp[k][:10]
	//	}
	//	for _, v2 := range mp[k] {
	//		cnt[v2.GroupId] += v2.Goal
	//	}
	//}
	//
	//slice := make([]int, 0, len(cnt))
	//for k, _ := range cnt {
	//	slice = append(slice, k)
	//}
	//sort.Slice(slice, func(a, b int) bool {
	//	return cnt[slice[a]] > cnt[slice[b]]
	//})
	//groupId := 0
	//if usr, ok := c.Get("user"); ok {
	//	groupId = usr.(model.User).GroupId
	//}
	//c.JSON(http.StatusOK, utils.H{
	//	"message": "success",
	//	"group":   groupId,
	//	"data": map[string]interface{}{
	//		"rank":     slice,
	//		"scoreSum": cnt,
	//		"detail":   mp,
	//	},
	//})
}
