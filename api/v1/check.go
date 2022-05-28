/**
  @author: 云坠  2021/12/28 20:34:00
  @note:
*/
package v1

import (
	"DuDao/dto"
	"DuDao/models"
	"DuDao/models/Res"
	"DuDao/models/common"
	"DuDao/models/request"
	"DuDao/models/response"
	"DuDao/services/system"
	"DuDao/util"
	"github.com/gin-gonic/gin"
)

//用户注册
func Register(c *gin.Context){
	var r request.Register
	_ = c.ShouldBindJSON(&r)
	uuid := util.GetUuid()
	r.Password = util.MD5V([]byte(r.Password))
	user:=models.User{
		Id:        uuid,
		Account:   r.Account,
		Password:  r.Password,
		UserName:      r.Username,
		Authority: "",
	}
	err := userservice.Register(user)
	if err != nil{
		response.FailWithMessage(err.Error(),c)
	}else {
		response.OkWithMessageAndDetail("",user,c)
	}
}
//LoginHandle 登陆
func LoginHandle(c *gin.Context) {
	request := request.LoginRequest{}

	err := c.ShouldBindJSON(&request)
	request.Password = util.MD5V([]byte(request.Password))
	if err != nil {
		c.JSON(401, &Res.LoginResult{
			ReturnUrl: "null",
			Token:     "null",
			Authority: "null",
			Result:    "null",
			Message:   "Object reference not set to an instance of an object",
			Code:      500,
		})
		return
	}

	userDto := dto.UserDto{
		Account:  request.Account,
		Password: request.Password,
	}

	service := system.CheckService{}
	token, err, authority := service.Login(userDto)

	if err != nil {
		c.JSON(401, &Res.LoginResult{
			ReturnUrl: "null",
			Token:     "null",
			Authority: "",
			Result:    "null",
			Message:   "Object reference not set to an instance of an object",
			Code:      500,
		})
		return
	}

	c.SetCookie("X-Token", token, 14400, "/", "", false, false)

	c.JSON(200, Res.LoginResult{
		ReturnUrl: "",
		Token:     token,
		Authority: authority,
		Result:    "",
		Message:   "操作成功",
		Code:      200,
	})
}

// GetUserNameHandle 根据token获取用户名
func GetUserNameHandle(c *gin.Context) {
	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}
	service := system.CheckService{}

	result, err := service.GetUserName(token)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.JSON(200, Res.ResponseIn{
		Result:  result,
		Message: "操作成功",
		Code:    200,
	})
}


// 根据Token注销登录

func LogOutHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	service := system.CheckService{}
	result, err := service.LogOut(token)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.SetCookie("X-Token", token, -1, "/", "", false, false)
	c.JSON(200, Res.ResponseOut{
		Result:  result,
		Message: "操作成功",
		Code:    200,
	})
}

