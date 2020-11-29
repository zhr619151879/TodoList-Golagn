# TodoList-Golang
A simple TodoList Web application by Golang(using framework gin&amp;gorm)





* go.mod 有所需要的go依赖



* 文件结构:

 ├── controller
 │   └── controller.go
 ├── dao
 │   └── mysql.go
 ├── go.mod
 ├── go.sum
 ├── main.go
 ├── model
 │   └── todo.go
 ├── router
 │   └── router.go
 ├── static
 │   ├── css
 │   │   ├── app.708ce172.css
 │   │   └── chunk-vendors.57db8905.css
 │   ├── fonts
 │   │   ├── element-icons.535877f5.woff
 │   │   └── element-icons.732389de.ttf
 │   └── js
 │       ├── app.82bdd871.js
 │       └── chunk-vendors.ddcb6f91.js
 └── templates
     ├── favicon.ico
     └── index.html



controller->路由控制函数

dao->数据访问层

model->数据模型层(由于逻辑简单, 与服务整合)

static->静态资源

router->路由

static->前端打包的静态资源(前端vue项目也整合在此仓库中)



* 运行效果:









### launch



* backend文件夹: 



`go run main.go` 



url : localhost:8080



* 前端项目: frontend文件夹



`npm install`



* 打包

`npm build`



* 运行

`npm run`



---

* 有待完善....

用户定制功能
