package reqModel

import "scnu_acm_rank/biz/model"

type CreateTeamReq struct {
	Name string `form:"name,required" json:"name"`
	Key  string `form:"key,required" json:"key"`
}

func (req *CreateTeamReq) GetModel() *model.Team {
	model := &model.Team{}
	model.Name = req.Name
	return model
}
