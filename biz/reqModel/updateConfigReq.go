package reqModel

import "scnu_acm_rank/biz/model"

type UpdateConfigReq struct {
	VjUserName string `form:"vjUserName,required" json:"vjUserName"`
	VjPassWord string `form:"vjPassWord,required" json:"vjPassWord"`
	VjCookie   string `form:"vjCookie,required" json:"vjCookie"`
}

func (req *UpdateConfigReq) Convert2DbModel() *model.Config {
	resp := &model.Config{}
	resp.VjUserName = req.VjUserName
	resp.VjPassWord = req.VjPassWord
	resp.VjCookie = req.VjCookie
	return resp
}
