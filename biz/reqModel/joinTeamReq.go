package reqModel

type JoinTeamReq struct {
	Key   string `form:"key,required" json:"key"`
	StuId int64  `form:"stu_id,required"`
}
