package froms

type BokeForm struct {
	CategoryId int32  `form:"categoryid" json:"categoryid" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
	Title      string `form:"title" json:"title" binding:"required"`
}
type BokeForm1 struct {
	CategoryId int32  `form:"categoryid" json:"categoryid" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
	Title      string `form:"title" json:"title" binding:"required"`
	Id         int32  `form:"id" json:"id" binding:"required"`
}
type CategoryFrom struct {
	Id   string `form:"id" json:"id" binding:"required,min=1"`
	Name string `form:"name" json:"name" binding:"required,min=2,max=100"`
}
type CategoryFrom1 struct {
	Id   int32  `form:"id" json:"id" binding:"required,min=1"`
	Name string `form:"name" json:"name" binding:"required,min=2,max=100"`
}
type CategoryFrom2 struct {
	Id int32 `form:"id" json:"id" binding:"required,min=1"`
}
