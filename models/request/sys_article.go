/**
* @Author: 云坠
* @Date: 2022/5/27 19:11
**/
package request

type ArticleDetail struct {
	Content string  `gorm:"column:content"`
	Title   string  `gorm:"column:title"`
	Labels	string	`gorm:"column:labels"`
}
type QueryLoad struct {
	Id string			//用户id
	Page int			//要查看的页数,默认为1
	Limit int			//每页最大展示数量
	Order int			//记录排序方式,可选0,1,2,默认为0随机序,1升序,2降序
}