/**
  @author: qianyi  2022/1/1 17:42:00
  @note:
*/
package system

import (
	"DuDao/dto"
	"DuDao/global"
	"DuDao/models"
	"DuDao/pkg/helper"
	"DuDao/repositories"
	"errors"
	"fmt"
)

type UserService struct {
}

func (user UserService) ChangePassWd(token string,pwd string) error{
	claim,err := helper.Decode(token,helper.Key)

	if err != nil {
		return errors.New("解析Token失败")
	}

	userDto := dto.UserDto{
		Account: claim.Account,
		Password: pwd,
	}

	err = repositories.UpdateUser(userDto)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (userservice *UserService) Register(u models.User)(err error){
	var user models.User
	err = global.GVA_DB.Where("account = ?",u.Account).First(&user).Error//不存在返回record not found 存在返回nil
	if err == nil {
		fmt.Println(err)
		return errors.New("注册失败,账号已存在")
	}
	err = global.GVA_DB.Where("username = ?",u.UserName).First(&user).Error
	if err == nil {
		fmt.Println(err)
		return errors.New("注册失败:用户名已注册,请重新输入用户名")
	}

	err = global.GVA_DB.Create(&u).Error
	return err
}
func (userservice *UserService) GetUserId (token string) (userid string,err error){
	userid,err = repositories.GetUserId(token)
	if err!=nil {
		return "" ,err
	}
	return  userid,err
}
