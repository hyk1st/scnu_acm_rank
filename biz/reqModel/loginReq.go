package reqModel

type LoginReq struct {
	StuId    int64  `form:"stu_id,required"`
	Password string `form:"password,required"`
}
