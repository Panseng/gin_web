package model

import "gin_web/utils/status_msg"

type Profile struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20)" json:"name"`
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	Qqchat string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat string `gorm:"type:varchar(100)" json:"wechat"`
	Weibo  string `gorm:"type:varchar(200)" json:"weibo"`
	Bili   string `gorm:"type:varchar(200)" json:"bili"`
	Email  string `gorm:"type:varchar(200)" json:"email"`
	Img    string `gorm:"type:varchar(200)" json:"img"`
	Avatar string `gorm:"type:varchar(200)" json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)" json:"icp_record"`
}

// GetProfile 获取个人信息设置
func GetProfile(id int) (Profile, int) {
	var profile Profile
	err = db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, status_msg.ERROR
	}
	return profile, status_msg.SUCCSE
}

// UpdateProfile 更新个人信息设置
func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return status_msg.ERROR
	}
	return status_msg.SUCCSE
}
