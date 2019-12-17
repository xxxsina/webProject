# golang 学习

## 先 git 项目
> git init

> git remote add origin https://github.com/xxxsina/webProject.git

> git add .

> git commit -m "message"

> git push -u origin master



> git pull origin master


## 编译的秘密

### go run 和 go build 和 go install 命令区别

* go run：

* 　　go run 编译并直接运行程序，它会产生一个临时文件（但实际不存在，也不会生成 .exe 文件），直接在命令行输出程序执行结果，方便用户调试。

* 　　注意点：需要在main包下执行go run ，否则如下
	
	> go run ./mytest/mytest.go


* go build：

*	go build 用于测试编译包，主要检查是否会有编译错误，如果是一个可执行文件的源码（即是 main 包），就会在当前目录直接生成一个可执行文件。


* go install：

*	go install 的作用有两步：

*	第一步是编译导入的包文件，所有导入的包文件编译完才会编译主程序；

*	第二步是将编译后生成的可执行文件放到 bin 目录下（$GOPATH/bin），编译后的包文件放到 pkg 目录下（$GOPATH/pkg）。（$GOPATH为Go的工作目录）

```
    # 试用了几种办法，还是这种靠谱，一句搞定win环境编译成linux  
    # GOOS=linux GOARCH=amd64 go build
    
    # 这个也可以，但是生成的bin文件在$GOPATH/bin里面
    # GOOS=linux GOARCH=amd64 go install
```

#### 三者区别：

* go run 和 go build 后面是直接加xxx.go ; 从文件编译

* 而go install 后面是直接加xxx（xxx为目录名） ； 从文件夹编译


## 其他

### 已完成 gin 、 xorm 、 mysql 、JWT 、redis（完成了基本的set get exist操作）
### 未完成 MongoDB

## 项目的秘密

### 目录结构

- webPoject
    - app
        - config
            - app.go    //gin启动，监听等
            - common.go //公共配置数据结构
            - router.go //路由配置
        - controllers
            1. 存放控制器
        - data
            1. 资源文件夹
        - views
            1. 模板
        - main.go //入口
    - bin
        * 存放编辑后的文件
    - com_party
        - config //存放公共配置文件，基础配置和数据库配置
        - helper //工具
        - libraries //库
        - middleware //中间件
        - models //存放数据映射文件
        - service //存放业务逻辑文件
    - data
        * 存放各种静态资源文件
    - framework
        - DB //存放数据驱动文件，目前只有mysql，后期再优化
    - templates
        1. 存放的用于反射的文件
        2. 反射：xorm.exe reverse mysql "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8" templates/goxorm
    - go.mod //库管理文件
    - README.md

## LICENSE

### 热跟新(热启动)

 ```
    # go get github.com/pilu/fresh
    # 进入项目目录，含有main.go文件夹
    # ../../../bin/fresh.exe  //这里需要根据具体fresh存放的位置来
 ```

### 数据加密

```
    # go get github.com/wumansgy/goEncrypt
    # 如果采用AES的CBC模式，密钥key的长度为16、24、32位
```

BSD License
[https://www.cnblogs.com/Paul-watermelon/articles/10842752.html](https://www.cnblogs.com/Paul-watermelon/articles/10842752.html)
