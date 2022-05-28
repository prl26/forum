/**
  @author: qianyi  2021/12/28 20:36:00
  @note:
*/
package v1

import (
	v1 "DuDao/api/v1"
	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine)  {
	apiGroup := engine.Group("/api")
	// 用户操作
	userRouter := apiGroup.Group("/user")
	userRouter.POST("/ChangePassWd",v1.ChangePassWdHandle)
	userRouter.GET("/Register",v1.Register)
}


