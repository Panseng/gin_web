# gin_web
以gin为后台框架，使用mysql作数据库。以vue为前端框架\
参考：[项目](https://github.com/wejectchen/Ginblog)

## 差异
文件命名：本项目文件和文件夹不同单词均使用小写单词，通过下划线连接\
包命名：将errorMsg变更为statusMsg\
部分依赖更新到最新版本，使用当前版本语境进行代码编写产生的差异\
model部分：mysql的使用和理解上的差异\

## 目录结构

```shell
├─  .gitignore
│  go.mod // 项目依赖
│  go.sum
│  LICENSE
│  main.go //主程序
│  README.md
│  tree.txt
│          
├─api         
├─config // 项目配置入口   
├─database  // 数据库备份文件（初始化）
├─log  // 项目日志
├─middleware  // 中间件
├─model // 数据模型层
├─routes
│      router.go // 路由入口    
├─static // 打包静态文件
│  ├─admin  // 后台管理页面 (已废弃，打包静态文件在web/admin/dist下)         
│  └─front  // 前端展示页面 (已废弃，打包静态文件在web/front/dist下)            
├─upload   
├─utils // 项目公用工具库
│  │  setting.go 
│  ├─errmsg   
│  └─validator         
└─web // 前端开发源码（VUECLI项目源文件)
    ├─admin             
    └─front
```
## 运行&&部署

1. 克隆项目

   ```shell
   git clone https://github.com/Panseng/gin_web.git
   ```

2. 转到下面文件夹下


	cd yourPath/ginbolg


3. 安装依赖

```
go mod tidy
```

4. 初始化项目配置config.ini

```ini
./config/config.ini

[server]
AppMode = debug # debug 开发模式，release 生产模式
HttpPort = :3000 # 项目端口
JwtKey = 89js82js72 #JWT密钥，随机字符串即可

[database]
Db = mysql #数据库类型，不能变更为其他形式
DbHost = 127.0.0.1 # 数据库地址
DbPort = 3306 # 数据库端口
DbUser = GinWeb # 数据库用户名
DbPassWord = admin123 # 数据库用户密码
DbName = GinWeb # 数据库名
```

5. 在database中将sql文件导入数据库

	推荐navicat或者其他sql管理工具导入

6. 启动项目

```shell
 go run main.go
```

此时，项目启动，你可以访问页面

```shell
首页
http://localhost
后台管理页面
http://localhost/admin
```

默认管理员:admin  密码:123456

## 实现功能

1.  简单的用户管理权限设置
2.  用户密码加密存储
3.  文章分类自定义
4.  列表分页\
~~5.  图片上传七牛云~~
6.  JWT 认证
7.  自定义日志功能
8.  跨域 cors 设置
9.  文章评论功能

## 技术栈

- golang
  - Gin web framework
  - gorm(v1 && v2)
  - jwt-go
  - scrypt
  - logrus
  - gin-contrib/cors
  - go-playground/validator/v10
  - go-ini
- JavaScript
  - vue
  - vue cli
  - vue router
  - ant design vue
  - vuetify
  - axios
  - tinymce
  - moment
- MySQL version:8.0.21

