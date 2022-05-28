/**
* @Author: 云坠
* @Date: 2022/5/28 9:28
**/
package v1

import (
	v1 "DuDao/api/v1"
	"github.com/gin-gonic/gin"
)

func ImgRouters(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	// Img及登录相关接口
	ImgGroup := apiGroup.Group("/Img")
	ImgGroup.POST("/uploadImg",v1.UploadOne)
	ImgGroup.POST("/ShowImg",v1.ShowImg)
}

