package reqModel

type UpdateContestResultReq struct {
	ContestId int    `form:"contest_id,required"`
	Json      string `form:"json,required"`
}
