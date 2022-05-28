/**
* @Author: 云坠
* @Date: 2022/5/27 20:04
**/
package system

import (
	"DuDao/global"
	"DuDao/models"
	"DuDao/models/request"
	"DuDao/repositories"
	"time"
)

type ArticleService struct {

}
func (a *ArticleService) GetArticleDetail(Id string)(detail request.ArticleDetail, err error){
	detail,err = repositories.GetArticleDetail(Id)
	return detail,err
}
func (a *ArticleService) UpdateArt(Id string,article request.ArticleDetail)error{
	art := models.Article{}
	err := global.GVA_DB.Model(&art).Where("Id = ?",Id).Updates(models.Article{
		UserId:        "",
		Id:            "",
		CreateAt:      time.Time{},
		Content:       article.Content,
		Title:         article.Title,
		Likes:         0,
		ReplyNumber:   0,
		ReadingVolume: 0,
		Labels:        article.Labels,
	}).Error
	return err
}
func (a *ArticleService) GetAllArticle (load request.QueryLoad)(data []models.Article ,err error){
	data,err = repositories.LoadAllArticle(load)
	return data,err
}
func (a *ArticleService) GetHotArt()(data []request.ArticleDetail,err error){
	data,err = repositories.GetHotArt()
	return data,err
}