// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/tools"
	"github.com/jinzhu/gorm"
)

const salt = "cloudy"

type UserService struct {
	DB *gorm.DB
	repositories.IUserRepository
}

func (service *UserService) ServiceGetSelf(id string) (user models.UserDetail, err error) {
	return service.FindUserByID(service.DB, id)
}
func (service *UserService) ServiceGetUserByID(id string) (user models.UserDetail, err error) {
	return service.FindUserByID(service.DB, id)
}
func (service *UserService) ServiceGetUserByUserName(username string) (user models.UserDetail, err error) {
	return service.FindUserByUserName(service.DB, username)
}
func (service *UserService) ServiceUserCreate(userCreate models.UserCreate) (user models.UserDetail, err error) {
	userCreate.ID = 0
	userCreate.Pwd = tools.SHA256(tools.StringsJoin(userCreate.Password, salt))
	return service.InsertUser(service.DB, userCreate)

}
func (service *UserService) ServiceCheckUser(username, pwd string) (user models.UserDetail, err error) {
	loginPwd := tools.SHA256(tools.StringsJoin(pwd, salt))
	user, err = service.FindUserByUserNameAndPwd(service.DB, username, loginPwd)
	return

}
