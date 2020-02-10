package models

import (
	"fg-admin/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jameskeane/bcrypt"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null VARCHAR(191)"`
	Username string `gorm:"type:varchar(32)"`
	Password string `gorm:"type:varchar(200)" json:"-"`
	Email    string `gorm:"type:varchar(32)"`
	Phone    string `gorm:"type:varchar(32)"`
	RoleID   uint
	Role     Role
}

func FindUserByName(username string) User {
	u := User{}
	if err := DB.Where("username = ?", username).First(&u).Error; err != nil {
		fmt.Printf("UserAdminCheckLoginErr:%s", err)
	}
	return u
}

// 通过 id 获取 user 记录
func GetUserById(id uint) User {
	user := new(User)
	user.ID = id

	if err := DB.Preload("Role").First(user).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
	}

	return *user
}

// 通过 username 获取 user 记录
func GetUserByUserName(username string) *User {
	user := &User{Username: username}

	if err := DB.Preload("Role").First(user).Error; err != nil {
		fmt.Printf("GetUserByUserNameErr:%s", err)
	}

	return user
}

// 通过 id 删除用户
func DeleteUserById(id uint) {
	u := new(User)
	u.ID = id

	if err := DB.Delete(u).Error; err != nil {
		fmt.Printf("DeleteUserByIdErr:%s", err)
	}
}

// 分页获取用户记录
func GetAllUsers(name, orderBy string, page, limit int) (users []*User) {
	if err := GetAll(name, orderBy, page, limit).Preload("Role").Find(&users).Error; err != nil {
		fmt.Printf("GetAllUserErr:%s", err)
	}
	return
}

// 总用户数
func GetUsersCount() int {
	var count int
	DB.Model(&User{}).Count(&count)
	return count
}

// 创建新用户
func CreateUser(aul *User) (user *User) {
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash(aul.Password, salt)

	user = new(User)
	user.Username = aul.Username
	user.Password = hash
	user.Name = aul.Name
	user.Email = aul.Email
	user.Phone = aul.Phone
	user.RoleID = aul.RoleID

	if err := DB.Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}

	return
}

// 更新
func UpdateUser(uj *User, id uint) *User {
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash(uj.Password, salt)

	user := new(User)
	user.ID = id
	uj.Password = hash

	if err := DB.Model(user).Updates(uj).Error; err != nil {
		fmt.Printf("UpdateUserErr:%s", err)
	}

	return user
}

// 校验登录信息
func CheckPwd(username, password string) (tokenString string, status bool, msg string) {
	user := FindUserByName(username)
	var err error
	if user.ID == 0 {
		msg = "用户不存在"
		return
	}

	if ok := bcrypt.Match(password, user.Password); !ok {
		msg = "用户名或密码错误"
		return
	}

	// 获取权限列表
	authList := GetPermsByRoleId(user.RoleID)
	var authStr string
	for _, v := range authList {
		authStr += v.AuthPath
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &config.JwtClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		config.UserInfo{Username: user.Username, Uid: user.ID, RoleId: user.RoleID, AuthStr: authStr},
	}
	tokenString, err = token.SignedString([]byte("secret"))

	if err != nil {
		msg = err.Error()
		return
	}

	status = true
	msg = "登陆成功"
	return
}

func CreateSystemAdmin(roleId uint) *User {

	aul := new(User)
	aul.Username = config.Conf.Get("test.LoginUserName").(string)
	aul.Password = config.Conf.Get("test.LoginPwd").(string)
	aul.Name = config.Conf.Get("test.LoginName").(string)
	aul.RoleID = roleId

	user := GetUserByUserName(aul.Username)

	if user.ID == 0 {
		fmt.Println("创建账号")
		return CreateUser(aul)
	} else {
		fmt.Println("重复初始化账号")
		return user
	}
}
