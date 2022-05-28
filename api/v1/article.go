/**
* @Author: 云坠
* @Date: 2022/5/26 17:44
**/
package v1

import (
	"DuDao/global"
	"DuDao/models"
	"DuDao/models/common"
	"DuDao/models/request"
	"DuDao/models/response"
	"DuDao/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)
//	添加文章
func AddArticle(c *gin.Context){
	token, err := common.IsTokenValid(c)	//检验用户的token是否有效
	if err != nil {
		c.String(401,err.Error())
		return
	}
	articleadd := request.ArticleDetail{}
	_ = c.ShouldBindJSON(&articleadd)
	Userid,err := userservice.GetUserId(token)
	if err!=nil {
		response.FailWithMessage(err.Error(),c)
		return
	}
	uuid := util.GetUuid()
	article := models.Article{
		UserId:        Userid,
		Id:            uuid,
		CreateAt:      time.Time{},
		Content:       articleadd.Content,
		Title:         articleadd.Title,
		Likes:         0,
		ReplyNumber:   0,
		ReadingVolume: 0,
		Labels:        articleadd.Labels,
	}
	global.GVA_DB.Create(&article)
	response.OkWithMessageAndDetail("添加成功",article,c)
	err = global.GVA_Redis.ZAdd("key",redis.Z{
		Score:  0,
		Member: uuid,
	}).Err()
	if err != nil{
		panic(err)
	}
}

//	删除文章
func DeleteArticle(c *gin.Context){
	_, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}
	article := models.Article{}
	Id := c.Query("Id")
	err = global.GVA_DB.Where("Id =?",Id).Delete(&article).Error
	if err != nil {
		response.FailWithMessage("删除错误",c)
	}else {
		response.OkWithMessage("删除成功",c)
	}
}
//	修改文章
func UpdateArticle(c *gin.Context){
	_, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}
	article := request.ArticleDetail{}
	_ = c.ShouldBindJSON(&article)
	Id := c.Query("Id")
	err = articleservice.UpdateArt(Id,article)
	articleDetail,_ := articleservice.GetArticleDetail(Id)
	if err != nil {
		response.FailWithMessage(err.Error(),c)
	}else {
		response.OkWithMessageAndDetail("更新成功",articleDetail,c)
	}
}
//	查看文章详情
func ReadArticle(c *gin.Context){
	_, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}
	Id := c.Query("Id")
	articleDetail,err := articleservice.GetArticleDetail(Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else{
		global.GVA_Redis.ZIncrBy("key",1,Id)
		fmt.Println()
		response.OkWithMessageAndDetail("查询成功",articleDetail,c)
	}
}
// 查看所有文章
func ReadAllArticle(c *gin.Context){
	//验证用户的token是否有效
	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}
	// 可以从前端请求获取分页的配置.也可以直接写死
	page := c.Query("page")
	p, _ := strconv.Atoi(page)
	limit := c.Query("limit")
	l, _ := strconv.Atoi(limit)
	order1 := c.Query("order")
	order, _ := strconv.Atoi(order1)
	id,_ := userservice.GetUserId(token)
	var load = request.QueryLoad{
		Id: id,
		Page:  p,
		Limit: l,
		Order: order,
	}

	allArticle,err := articleservice.GetAllArticle(load)
	if err != nil{
		response.FailWithMessage("查询出错",c)
		return
	}else {
		c.JSON(200,allArticle)
	}
}
func GetHotArt (c *gin.Context) {
	_, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}
	hotArt,err := articleservice.GetHotArt()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}else {
		c.JSON(200,hotArt)
	}
}