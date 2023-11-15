package reqModel

import "scnu_acm_rank/biz/model"

type UserEditReq struct {
	VjName     string `form:"vj_name,required"`
	CfId       string `form:"cf_id,required"`
	StuId      int64  `form:"stu_id,required"`
	Grade      string `form:"grade,required"`
	Status     int    `form:"status"`
	Connection string `form:"connection"`
}

func (user *UserEditReq) Change2UserModel(m *model.User) {
	m.CfId = user.CfId
	m.VjName = user.VjName
	m.Grade = user.Grade
	m.Status = user.Status
	m.Connection = user.Connection
}
