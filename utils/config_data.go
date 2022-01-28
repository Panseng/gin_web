package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strconv"
)
var (
	ServerMode	string
	ServerPort		string
	JwtKey   string

	DB			string
	DBHost		string
	DBPort		string
	DBUser		string
	DBPwd		string
	DBName	string

	AK string
	SK string
	Bucket string
	QiNiuServer string

	UserRoleManager int
	UserRoleGeneral int
)

func init()  {
	file, err := ini.Load("config/config.ini")
	if err != nil{
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	// 获取配置数据
	loadServer(file)
	loadDB(file)
	loadQiNiu(file)
	loadEnumerate(file) // 枚举常数
}

func loadServer(file *ini.File)  {
	// 获取[server]篇段，key值ServerMode对于的值，默认值为 debug
	ServerMode =file.Section("server").Key("ServerMode").MustString("debug")
	ServerPort =file.Section("server").Key("ServerPort").MustString("80")
}

func loadDB(file *ini.File)  {
	DB =file.Section("database").Key("DB").MustString("mysql")
	DBHost =file.Section("database").Key("DBHost").MustString("localhost")
	DBPort =file.Section("database").Key("DBPort").MustString("3306")
	DBUser =file.Section("database").Key("DBUser").MustString("root")
	DBPwd =file.Section("database").Key("DBPwd").MustString("mysql@2020")
	DBName =file.Section("database").Key("DBName").MustString("gin_web")
}

func loadQiNiu(file *ini.File)  {
	AK =file.Section("qi_niu").Key("AK").String()
	SK =file.Section("qi_niu").Key("SK").String()
	Bucket =file.Section("qi_niu").Key("Bucket").String()
	QiNiuServer =file.Section("qi_niu").Key("QiNiuServer").String()
}

func loadEnumerate(file *ini.File)  {
	UserRoleManager,_ =strconv.Atoi(file.Section("qi_niu").Key("AK").String())
	UserRoleGeneral,_ =strconv.Atoi(file.Section("qi_niu").Key("AK").String())
}