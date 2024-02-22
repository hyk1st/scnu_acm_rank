package reqModel

type SendEmailReq struct {
	Email string `form:"email,required" vd:"email($)"`
	StuId int64  `form:"stu_id"`
}
