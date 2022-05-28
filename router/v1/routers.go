/**
  @author: qianyi  2022/3/1 19:08:00
  @note:
*/
package v1

import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = [] Option{}

// 注册app的路由配置


func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		c.Header("Access-Control-Allow-Origin", "*")   // 可将将 * 替换为指定的域名
//		c.Header("Access-Control-Allow-Headers", "Content-TypeAccessToken,X-CSRF-Token, Authorization,x-token,content-type")    //你想放行的header也可以在后面自行添加
//		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")      //我自己只使用 get post 所以只放行它
//		//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,x-token,X-Token")
//		c.Header("Access-Control-Allow-Credentials", "true")
//
//		// 放行所有OPTIONS方法
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//		// 处理请求
//		c.Next()
//	}
//}

func RInit() *gin.Engine {
	r := gin.Default()
	//r.Use(middle.Cors())
	for _, opt := range options {
		opt(r)
	}
	return r
}