package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Permission struct {
	gorm.Model
	Name        string `gorm:"unique;not null VARCHAR(191)" json:"name" validate:"required,gte=4,lte=50"`
	Description string `gorm:"VARCHAR(191)" json:"description"`
	AuthPath    string `gorm:"VARCHAR(20)"`
	Pid         int    `gorm:"smallint"`
	Icon        string `gorm:"VARCHAR(20)"`
	IsShow      int    `gorm:"smallint"`
	Sort        int    `gorm:"smallint"`
}

// 通过 id 获取 permission 记录
func GetPermissionById(id uint) *Permission {
	permission := new(Permission)
	permission.ID = id

	if err := DB.First(permission).Error; err != nil {
		fmt.Printf("GetPermissionByIdError:%s", err)
	}

	return permission
}

// 通过 name 获取 permission 记录
func GetPermissionByName(name string) *Permission {
	permission := new(Permission)
	permission.Name = name

	if err := DB.First(permission).Error; err != nil {
		fmt.Printf("GetPermissionByNameError:%s", err)
	}

	return permission
}

// 通过ID删除权限
func DeletePermissionById(id uint) {
	u := new(Permission)
	u.ID = id

	if err := DB.Delete(u).Error; err != nil {
		fmt.Printf("DeletePermissionByIdError:%s", err)
	}
}

// 获取所有权限
func GetAllPermissions(name, orderBy string, offset, limit int) (permissions []*Permission) {
	if err := GetAll(name, orderBy, offset, limit).Find(&permissions).Error; err != nil {
		fmt.Printf("GetAllPermissionsError:%s", err)
	}

	return
}

// 创建新权限
func CreatePermission(aul *Permission) (permission *Permission) {
	if err := DB.Create(aul).Error; err != nil {
		fmt.Printf("CreatePermissionError:%s", err)
	}
	return
}

// 更新权限
func UpdatePermission(pj *Permission, id uint) (permission *Permission) {
	permission = new(Permission)
	permission.ID = id

	if err := DB.Model(&permission).Updates(pj).Error; err != nil {
		fmt.Printf("UpdatePermissionError:%s", err)
	}

	return
}

// 创建系统管理员
func CreateSystemAdminPermission() *Permission {
	aul := new(Permission)
	aul.Name = "/account/create"
	aul.Description = "创建账号权限"

	permission := GetPermissionByName(aul.Name)

	if permission.ID == 0 {
		fmt.Println("创建账号权限")
		return CreatePermission(aul)
	} else {
		fmt.Println("重复初始化权限")
		return permission
	}
}
