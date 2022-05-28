package request

//注册入参
type Register struct {
	Username string `json:"username"`
	Account string `json:"account"`
	Password string `json:"password"`
}
//LoginRequest 登陆入参
type LoginRequest struct {
	Account string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ChangePwdRep 修改密码入参
type  ChangePwdRep struct {
	Password string `form:"passwd" json:"passwd" binding:"required"`
}