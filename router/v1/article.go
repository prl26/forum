/**
* @Author: 云坠
* @Date: 2022/5/26 17:47
**/
package v1

import (
	v1 "DuDao/api/v1"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(engine *gin.Engine)  {
	apiGroup := engine.Group("/api")
	// 用户操作
	articleRouter := apiGroup.Group("/articles")
	articleRouter.GET("/Add",v1.AddArticle)
	articleRouter.GET("/Delete",v1.DeleteArticle)
	articleRouter.GET("/Change",v1.UpdateArticle)
	articleRouter.POST("/Read",v1.ReadArticle)
	articleRouter.POST("/ReadAll",v1.ReadAllArticle)
	articleRouter.POST("/GetHot",v1.GetHotArt)
}

