/**
* @Author: 云坠
* @Date: 2022/5/27 21:01
**/
package v1

import (
	"DuDao/global"
	"DuDao/models"
	"DuDao/models/response"
	"DuDao/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path"
	"strings"
)

//单张图片上传
func UploadOne(c *gin.Context){
	//获取表单数据 参数为name值
	f, err := c.FormFile("f1")
	//错误处理
	if err != nil {
		response.FailWithMessage("上传错误",c)
		return
	} else {
		//判断文件格式
		fileExt:=strings.ToLower(path.Ext(f.Filename))
		if fileExt!=".png"&&fileExt!=".jpg"&&fileExt!=".gif"&&fileExt!=".jpeg"{
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
			})
			return
		}
		//将文件保存至本项目static文件夹的views下
		uuid := util.GetUuid()
		dst := path.Join("./static/views",uuid+fileExt)
		c.SaveUploadedFile(f, dst)
		i := models.Img{
			Id:  uuid,
			Url: dst,
		}
		global.GVA_DB.Create(&i)
		//保存成功返回正确的Json数据
		response.OkWithMessage("保存成功",c)
	}
}
func ShowImg (c *gin.Context){
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
	}
	image,_ := ioutil.ReadFile("./static/views/ccd6154ee34c4e3499cfc2571b69bf0f.png") //暂时只写到了展示一张固定图片
	c.Writer.WriteString(string(image))
}



