package froms

type PassWordLoginForm struct {
	NickName string `form:"nickname" json:"nickname" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}

type RegisterForm struct {
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	NickName string `form:"NickName" json:"NickName" binding:"required"`
}
