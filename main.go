/**
  @author: 云坠  2021/12/30 22:08:00
  @note:
*/
package main

import (
	"DuDao/global"
	"DuDao/initialize"
	"DuDao/middle"
	v1 "DuDao/router/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	global.GVA_DB = initialize.GormMysql()
	global.GVA_LOG = middle.Zap()       // 初始化zap日志库
	global.GVA_Redis = initialize.NewRedisHelper()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	v1.Include(v1.CheckRouters,v1.UserRouter,v1.ImgRouters,v1.ArticleRouter)
	engine := v1.RInit()
	engine.LoadHTMLFiles("static/index.html")
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//err := global.GVA_Redis.ZAdd("key",redis.Z{
	//	Score:  0,
	//	Member: 1,
	//}).Err()
	//if err != nil{
	//	panic(err)
	//}
	//err = global.GVA_Redis.ZAdd("key",redis.Z{
	//	Score:  1,
	//	Member: 2,
	//}).Err()
	//if err != nil{
	//	panic(err)
	//}
	//rk, _ := global.GVA_Redis.ZRank("key", "2").Result()
	//fmt.Println(rk)
	//fmt.Println("sadfa")
	engine.Run()
}


