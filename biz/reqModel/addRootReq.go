package reqModel

type AddRootReq struct {
	StuID int64 `form:"stu_id,required"`
	Level int   `form:"level,required"`
}
