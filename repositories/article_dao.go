/**
* @Author: 云坠
* @Date: 2022/5/27 20:07
**/
package repositories

import (
	"DuDao/global"
	"DuDao/models"
	"DuDao/models/request"
	"fmt"
	"github.com/go-redis/redis"
)

func GetArticleDetail(Id string) (detail request.ArticleDetail, err error) {

	var article models.Article
	err = global.GVA_DB.Table("article").Where("Id = ?", Id).Find(&article).Error
	detail = request.ArticleDetail{
		Content: article.Content,
		Title:   article.Title,
		Labels:  article.Labels,
	}
	global.GVA_DB.Table("article").Where("id = ?",article.Id).Update("readingVolume",article.ReadingVolume+1)
	return detail, err
}
func LoadAllArticle(load request.QueryLoad) (data []models.Article, err error) {
	var count int64
	var pageindex,pagesize int
	// 分页功能
	if load.Limit <= 0 {
		load.Limit = 10
		pagesize = load.Limit
	} else {
		pagesize = load.Limit
	}
	if load.Page <= 0 {
		pageindex = 1
	} else {
		pageindex = load.Page
	}
	if load.Order == 0 {
		global.GVA_DB.Order("rand()")
	}

	if load.Order == 1 {
		global.GVA_DB.Order("addAt ASC")

	}
	if load.Order == 2 {
		global.GVA_DB.Order("addAt DESC")
	}
	global.GVA_DB.Count(&count).Find(&data)
	global.GVA_DB.Offset((pageindex - 1) * pagesize).Limit(pagesize).Where("userId =?",load.Id).Find(&data)
	fmt.Printf("共查询到%d条数据",count)
	return data,nil
}
func GetHotArt ()(data []request.ArticleDetail, err error){
	var hotArtnumber []string
	//此处必须定义切片长度,否则会报错
	var hotArt = make([]request.ArticleDetail,10)
	//var hotArt  []request.ArticleDetail
	op := redis.ZRangeBy{
		Min:    "0",
		Max:    "100000000",
		Offset: 0,
		Count:  10,
	}
	vals, err := global.GVA_Redis.ZRevRangeByScore("key",op).Result()			//对集合中的数据根据score从大到小排序
	if err != nil {
		panic(err)
	}
	for _, val := range vals {
		hotArtnumber = append(hotArtnumber, val)
		//fmt.Println(vals)
	}
	//fmt.Println(hotArtnumber)
	for i := 0; i<len(hotArtnumber); i++{
		var hotArt_i request.ArticleDetail
		global.GVA_DB.Table("article").Where("id = ?",hotArtnumber[i]).Find(&hotArt_i)
		fmt.Println(hotArt_i)
		hotArt[i] = hotArt_i
	}
	return hotArt,err
}
