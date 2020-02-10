package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name        string        `gorm:"unique;not null VARCHAR(191)"`
	Description string        `gorm:"VARCHAR(191)"`
	Perms       []*Permission `gorm:"many2many:role_perms;"`
}

func GetPermsByRoleId(id uint) []Permission {
	var permission []Permission
	role := Role{}
	DB.First(&role, "id = ?", id)
	DB.Model(&role).Related(&permission,  "Perms")
	return  permission
}

// 通过 id 获取 role 记录
func GetRoleById(id uint) *Role {
	role := new(Role)
	role.ID = id

	if err := DB.Preload("Perms").First(role).Error; err != nil {
		fmt.Printf("GetRoleByIdErr:%s", err)
	}

	return role
}

// 通过 name 获取 role 记录
func GetRoleByName(name string) *Role {
	role := new(Role)
	role.Name = name

	if err := DB.Preload("Perms").First(role).Error; err != nil {
		fmt.Printf("GetRoleByNameErr:%s", err)
	}

	return role
}

// 通过 id 删除角色
func DeleteRoleById(id uint) {
	u := new(Role)
	u.ID = id

	if err := DB.Delete(u).Error; err != nil {
		fmt.Printf("DeleteRoleErr:%s", err)
	}
}

// 获取所有的角色
func GetAllRoles(name, orderBy string, offset, limit int) (roles []Role) {

	if err := GetAll(name, orderBy, offset, limit).Preload("Perms").Find(&roles).Error; err != nil {
		fmt.Printf("GetAllRoleErr:%s", err)
	}
	return
}

// 创建角色
func CreateRole(aul *Role, permIds []uint) (role *Role) {

	role = new(Role)
	role.Name = aul.Name
	role.Description = aul.Description

	if err := DB.Create(role).Error; err != nil {
		fmt.Printf("CreateRoleErr:%s", err)
	}

	perms := []Permission{}
	DB.Where("id in (?)", permIds).Find(&perms)
	if err := DB.Model(&role).Association("Perms").Append(perms).Error; err != nil {
		fmt.Printf("AppendPermsErr:%s", err)
	}

	return
}


func UpdateRole(rj *Role, id uint, permIds []uint) (role *Role) {
	role = new(Role)
	role.ID = id

	if err := DB.Model(&role).Updates(rj).Error; err != nil {
		fmt.Printf("UpdateRoleErr:%s", err)
	}

	perms := []Permission{}
	DB.Where("id in (?)", permIds).Find(&perms)
	if err := DB.Model(&role).Association("Perms").Replace(perms).Error; err != nil {
		fmt.Printf("AppendPermsErr:%s", err)
	}

	return
}

// 创建系统管理员
func CreateSystemAdminRole(permIds []uint) *Role {
	aul := new(Role)
	aul.Name = "admin"
	aul.Description = "超级管理员"

	role := GetRoleByName(aul.Name)

	if role.ID == 0 {
		fmt.Println("创建角色")
		return CreateRole(aul, permIds)
	} else {
		fmt.Println("重复初始化角色")
		return role
	}
}
