package reqModel

import (
	"scnu_acm_rank/biz/model"
	"time"
)

type CreateContestReq struct {
	Name      string    `form:"name, required" json:"name"`
	VjCpId    string    `form:"vj_cp_id, required" json:"vj_cp_id"`
	Kind      int       `form:"kind, required" json:"kind"`
	Password  string    `form:"password, required" json:"password"`
	StartDate time.Time `form:"start_date, required" json:"start_date"`
	Length    int64     `form:"length" json:"length"`
	Result    string    `form:"result, required" json:"result"`
	CreateUsr string    `form:"create_user, required" json:"create_usr"`
	Ext       string    `form:"ext, required" json:"ext"`
}

func (req *CreateContestReq) Convert2model() *model.Competiton {
	resp := model.Competiton{}
	resp.Ext = req.Ext
	resp.Name = req.Name
	resp.StartDate = req.StartDate
	resp.Kind = req.Kind
	resp.Password = req.Password
	resp.CreateUsr = req.CreateUsr
	resp.Length = req.Length
	resp.VjCpId = req.VjCpId
	return &resp
}
