package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
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
)

func init()  {
	file, err := ini.Load("config/config.ini")
	if err != nil{
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	// 获取配置数据
	loadServer(file)
	loadDB(file)
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