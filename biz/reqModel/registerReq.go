package reqModel

import (
	"scnu_acm_rank/biz/model"
)

type RegisterReq struct {
	Email      string `form:"email,required" vd:"email($)"`
	Password   string `form:"password,required"`
	VjName     string `form:"vj_name,required"`
	CfId       string `form:"cf_id,required"`
	StuId      int64  `form:"stu_id,required"`
	Name       string `form:"name,required"`
	Sex        int    `form:"sex,required"`
	GroupId    int    `form:"group_id,required"`
	Grade      string `form:"grade,required"`
	Status     int    `form:"status"`
	Connection string `form:"connection"`
	Ext        string `form:"ext"`
}

func (r *RegisterReq) ToDbModle() *model.User {
	return &model.User{
		Password:   r.Password,
		VjName:     r.VjName,
		CfId:       r.CfId,
		StuId:      r.StuId,
		Name:       r.Name,
		Sex:        r.Sex,
		GroupId:    r.GroupId,
		Grade:      r.Grade,
		Connection: r.Connection,
		Ext:        r.Ext,
	}
}
func (r *RegisterReq) RegisterCheck(s *[]string) {
	if r.Sex < 0 || r.Sex > 1 {
		*s = append(*s, "性别参数错误")
	}
}
