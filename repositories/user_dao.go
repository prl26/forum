/**
  @author: qianyi  2021/12/28 20:14:00
  @note:
*/
package repositories

import (
	"DuDao/dto"
	"DuDao/global"
	"DuDao/models"
	"DuDao/pkg/helper"
	"errors"
	"fmt"
	"log"
)


type Addby struct {
	Addby string `gorm:"column:addBy"`
}

// GetUserByAccount  通过账户查询用户
func GetUserByAccount(account string) models.User {
	user := models.User{}
	global.GVA_DB.Find(&user, global.GVA_DB.Where("account = ?",account))
	return user
}
//UpdateUser 修改用户密码
func UpdateUser(userDto dto.UserDto) error {
	user := models.User{}

	if userDto.Password=="" {
		return errors.New("密码不能为空")
	}

	err := global.GVA_DB.Model(&user).Where("mobilePhone = ?",userDto.Account).Update("password",userDto.Password).Error

	return err
}

func GetUserByAddby(addby string) models.User {
	user := models.User{}
	global.GVA_DB.Find(&user, global.GVA_DB.Where("id = ?", addby))
	return user
}

func GetAddbyByDataID(dataID string)  string {
	var addby Addby
	global.GVA_DB.Table("tdata").Find(&addby, global.GVA_DB.Where("id = ?",dataID))
	return addby.Addby
}
func GetUserId(token string)(Id string,err error){
	claim, err := helper.Decode(token, helper.Key)
	if err != nil {
		log.Println("解析token失败")
		return "", errors.New("解析失败，请重新登录")
		fmt.Println("解析失败")
	}
	 var user models.User
	 global.GVA_DB.Where("account = ?",claim.Account).Find(&user)
	return user.Id,nil
}











////CreateUser 创建用户
//func CreateUser(dto dto.UserDto) error {
//	system := models.User{}
//	system.Account = dto.Account
//	system.Password = dto.Password
//
//	err := global.GVA_DB.Create(&system).Error
//
//	return err
//}
//
//// 修改密码
//func ChangePassword( rep v1.ChangePwdRep) error {
//	system := models.User{}
//
//	if rep.Password != "" {
//		system.Password = rep.Password
//	}
//
//	err := global.GVA_DB.Model(&system).Updates(&system).Error
//
//	return err
//}





