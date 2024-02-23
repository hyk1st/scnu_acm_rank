package reqModel

type CreateContestReq struct {
	CpId string `form:"cp_id,required" json:"cp_id"`
	Kind int    `form:"kind,required" json:"kind"`
	Name string `form:"name,required" json:"name"`
	//StartDate int64  `form:"start_date,required" json:"start_date"`
	//Length    int64  `form:"length"`
	//Result    string `form:"result"`
	//Json string `form:"json"`
	//Ext  string `form:"ext" json:"ext"`
}

//func (req *CreateContestReq) Convert2model() *model.Competition {
//	resp := model.Competition{}
//	resp.Ext = req.Ext
//	//resp.Name = req.Name
//	resp.Kind = req.Kind
//	//resp.Password = req.Password
//	//resp.CreateUsr = req.CreateUsr
//	//resp.Length = req.Length
//	//resp.VjCpId = req.VjCpId
//	return &resp
//}
