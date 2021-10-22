# go-study 一阶段实践项目博客系统

## 功能
* 使用 gin 作为基础web服务框架
* 使用 JWT (JsonWebToken)
* 使用 validator.v10 作为表单验证组件
* 使用 viper 管理配置文件

## 目录结构
```code 
├── config		    数据库和服务器相关配置信息			
├── internal      	    私有应用目录
│   ├── app
│   │   └── controller      控制器
│   │        └── v1   	    v1相关控制器
│   ├── pkg		    自定义包文件
│   │     ├── common        分页信息和公共参数
│   │     ├── config        数据库和服务器相关配置加载
│   │     ├── core          响应
│   │     └── middleware    中间件
│   ├── models		    数据模型
│   ├── routes       	    路由文件
│   └── service      	    逻辑服务层
├── main.go                 启动入口文件
├── go.mod
├── go.sum
└── blog                    数据库表信息 
```
### 导入blog sql数据库文件

## 启动
- cd /项目根路径
- go run main.go