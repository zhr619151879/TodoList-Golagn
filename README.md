# TodoList-Golang
A simple TodoList Web application by Golang(using framework gin&amp;gorm)





* go.mod 有所需要的go依赖



* 文件结构:

![image](https://github.com/zhr619151879/TodoList-Golang/png/tree.jpg)


controller->路由控制函数

dao->数据访问层

model->数据模型层(由于逻辑简单, 与服务整合)

static->静态资源

router->路由

static->前端打包的静态资源(前端vue项目也整合在此仓库中)



* 运行效果:

![image](https://github.com/zhr619151879/TodoList-Golang/png/image-20201129202651763.png)






### launch



* backend文件夹: (前端项目已打包静态入此文件夹, 直接运行即可访问)



`go run main.go` 



url : localhost:8080



* 前端项目: frontend文件夹(若要调试)



`npm run serve`



* 打包

`npm run build`



* 运行

`npm run`



---

* 有待完善....

用户定制功能
