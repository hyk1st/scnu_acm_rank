package reqModel

import "scnu_acm_rank/biz/model"

type CreateTeamReq struct {
	Name string `form:"name,required" json:"name"`
}

func (req *CreateTeamReq) GetModel() *model.Team {
	model := &model.Team{}
	model.Name = req.Name
	return model
}
