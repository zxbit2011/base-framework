// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type IRoleRepository interface {
	//根据ID查找
	FindRoleByID(DB *gorm.DB, id string) (role models.RoleDetail, err error)
	// 创建角色
	InsertRole(DB *gorm.DB, role models.RoleCreate) (result models.RoleDetail, err error)
	// 修改角色
	UpdateRole(DB *gorm.DB, role models.RoleUpdate) (result models.RoleDetail, err error)
	// 获得所有角色
	FindAllRole(DB *gorm.DB) (roles []*models.RoleList, err error)
}
