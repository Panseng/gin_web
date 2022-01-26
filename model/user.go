package model

import (
	"gin_web/utils/status_msg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	// 1为管理员，2为普通用户
	Role int `gorm:"type:int;DEFAULT:2" json:"role" label:"角色码"`
}

// CheckUserName 查询用户是否存在
func CheckUserName(name string) int  {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0{
		return statusMsg.ERROR_USERNAME_USED // 1001
	}
	return statusMsg.SUCCSE // 200
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	err :=db.Create(&data).Error
	if err != nil{
		return statusMsg.ERROR // 500
	}
	return statusMsg.SUCCSE // 200
}

// GetUser 查询用户
func GetUser(id int) (User, int)  {
	var user User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, statusMsg.ERROR
	}
	return user, statusMsg.SUCCSE
}

// GetUserList 获取用户列表
func GetUserList(username string, pageSize, pageNum int)  ([]User, int64){
	var userList []User
	var total int64
	if username != ""{
		db.Select("id,username,role,created_at").Where(
			"username LIKE ?", "%"+username+"%", // 含有该内容的都查询出来
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userList)
		db.Model(&userList).Count(&total)
		return userList, total
	}
	db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userList)
	db.Model(&userList).Count(&total)
	return userList, total
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) int  {
	var user User
	// 多条件查询，如果参考源代码，则会出错，因为查询出来的 username可能是其他人的历史名称
	db.Select("id, username").Where("id = ? and username = ?", id, name).First(&user)
	if user.ID > 0 {
		return statusMsg.SUCCSE
	}
	return statusMsg.ERROR
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	//var user User
	//var maps = make(map[string]interface{})
	// 通过map来匹配转换修改的值
	//maps["username"] = data.Username
	//maps["role"] = data.Role
	//err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error

	err = db.Select("username, role").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return statusMsg.ERROR
	}
	return statusMsg.SUCCSE
}

// ChangePassword 修改密码
func ChangePassword(id int, data *User) int {
	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return statusMsg.ERROR
	}
	return statusMsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return statusMsg.ERROR
	}
	return statusMsg.SUCCSE
}

// BeforeCreate gorm提供的hook 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	if u.Role != 1{ // 1: 管理员, 2: 普通用户
		u.Role = 2
	}
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// ScryptPw 生成密码 hash 值
func ScryptPw(password string) string {
	const cost = 10 // 4-31，值越高开销越大，破解难度越大

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	if user.ID == 0 { // 用户不存在
		return user, statusMsg.ERROR_USER_NOT_EXIST
	}
	if user.Role != 1 { // 非管理员
		return user, statusMsg.ERROR_USER_NO_RIGHT
	}

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if PasswordErr != nil {
		return user, statusMsg.ERROR_PASSWORD_WRONG
	}
	return user, statusMsg.SUCCSE
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)
		if user.ID == 0 { // 用户不存在
		return user, statusMsg.ERROR_USER_NOT_EXIST
	}

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if PasswordErr != nil {
		return user, statusMsg.ERROR_PASSWORD_WRONG
	}
	return user, statusMsg.SUCCSE
}