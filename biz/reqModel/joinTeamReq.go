package reqModel

type JoinTeamReq struct {
	Key string `form:"key,required" json:"key"`
}
