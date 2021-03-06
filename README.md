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

* https://www.cnblogs.com/Paul-watermelon/articles/10842752.html

## 其他

### 已完成 gin 、 xorm 、 mysql 、JWT 、redis（完成了基本的set get exist操作）
### 未完成 MongoDB、NSQ

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
        ```
        > 1 操作都是全局的
        # cd ($GOPATH/src)/golang.org/x
        # git clone https://github.com/golang/crypto.git
        > 2
        # cd ($GOPATH)/src/cloud.google.com/
        # git clone  https://github.com/googleapis/google-cloud-go.git  //这里不容易下下来，就去网站直接下win版的
        > 3
        # go get github.com/go-xorm/cmd/xorm
        # go install github.com/go-xorm/cmd/xorm    //这时候在bin里面生成了一个xorm.exe文件
        >> 4 开始执行reverse了，在bin目录下
        # ./xorm.exe reverse mysql "root:root@tcp(127.0.0.1:3306)/[databaseName]?charset=utf8" ../src/templates/goxorm ../src/models
        >>> 5 生成的models也在bin目录下
        ```
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

### 其他

1.修改代理
```
    # 说明：https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md
    # Go 1.13 及以上
    # go env -w GOPROXY=https://goproxy.cn,direct
```

### NSQ 安装
1、下载地址
```
    # https://nsq.io/deployment/installing.html
    sudo wget nsq-1.2.0.linux-amd64.go1.12.9.tar.gz
```
2、解压 
```
    tar -zxvf nsq-1.2.0.linux-amd64.go1.12.9.tar.gz
```
3、添加环境变量 
```
    sudo vim /etc/profile.d/nsq.sh
        # 加入以下
        # export PATH=$PATH:/home/[your-path]/bin
    # 刷新配置 
    source /etc/profile
```
4、NSQadmin地址
```
    http://nsqadmin.xxx.com/
```
5、配置简单集群start和end
```
    # 启动
    ../../nsq/sh/nsqstart.sh
    # 停止
    ../../nsq/sh/nsqend.sh
```

### MongoDB 安装

1、直接去官网看
```
    https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/
```
2、查看版本的方法
```
    # mongo
    # 或
    # mongod --version
    # 或
    # mongo --version
    # 或
    # mongod
```
3、配置文件mongod.conf所在路径
```
    # cd /etc/mongod.conf   #配置文件
    
    # dbPath: /var/lib/mongodb   #数据库存储路径
    # logAppend: true     #以追加的方式写入日志
    # path: /var/log/mongodb/mongod.log   #日志文件路径
    # port: 27017 #端口
    # bindIp: 127.0.0.1   #绑定监听的ip 127.0.0.1只能监听本地的连接，可以改为0.0.0.0

    # auth=true #设置需要登录权限
    ## 进入mongo，输入如下命令设置管理员账号密码
    # mongo --port 27017
    # 进入数据库，并创建用户和设置密码及给定权限(userAdminAnyDatabase[创建和管理数据库用户的权限]、readWrite[读取和写入数据库])
    > use admin;
    > db.createUser(
            {
                user: "root",
                pwd: "1qa2wsqazwsx",
                roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
            }
        )
    # 修改密码的命令如下
    > db.changeUserPassword('myUserAdmin','tUDfqjDWHR4hSIXs')
    # 验证密码是否添加成功
    > db.auth("root", "1qa2wsqazwsx") #如果返回1，则表示成功。
    # 要使密码生效，配置文件里面：找到#security: 取消注释
    > security:
        > authorization: enabled #注意缩进，缩进参照配置文件其他配置。缩进错误可能重启不成功。
    ## 连接数据库，并创建用户
    # mongo --port 27017 -u "myUserAdmin" -p "tUDfqjDWHR4hSIXs" --authenticationDatabase "admin"
    > use bili
    > db.createUser(
            {
                user: "moji_01",
                pwd: "1qaz2wsx",
                roles: [ { role: "readWrite", db: "moji" } ]
            }
        )
    > db.auth("moji_01", "1qaz2wsx")

    ## 以test用户连接即可操作数据库bili
    # mongo --port 21071 -u "moji_01" -p "1qaz2wsx" --authenticationDatabase "moji"
    > use moji
    # 这里需要默认添加一条数据进去，要不然就是坑啊
    > db.config.insert({"databaseName":"moji", "object":"moji"})
    WriteResult({ "nInserted" : 1 })

    ps: 添加用户时各个角色对应权限
    1.数据库用户角色：read、readWrite;
    2.数据库管理角色：dbAdmin、dbOwner、userAdmin；
    3.集群管理角色：clusterAdmin、clusterManager、clusterMonitor、hostManager；
    4.备份恢复角色：backup、restore
    5.所有数据库角色：readAnyDatabase、readWriteAnyDatabase、userAdminAnyDatabase、dbAdminAnyDatabase
    6.超级用户角色：root
    # https://segmentfault.com/a/1190000015603831
```
```
    # 更多相关操作
    # https://blog.csdn.net/zhangpeterx/article/details/88857699
    # https://www.cnblogs.com/cmyxn/p/6610297.html
    # https://blog.csdn.net/u010649766/article/details/79817549
    # https://blog.csdn.net/qq_38363459/article/details/80159387 #带密码的链接操作
    # https://cardinfolink.github.io/2017/05/17/mgo-session/ #连接池
    # http://www.jyguagua.com/?p=3126
    # https://blog.csdn.net/LK_whq/article/details/92414353 #增删改查
    # https://www.jianshu.com/p/fbe25ea58384 #项目参考
```
4、启动、关闭
```
    # sudo service mongod start  #启动
    # sudo service mongod stop   #关闭
    # ps aux | grep mongod   #查看守护进程mongod的运行状态
    
```
5、卸载
```
    # sudo service mongod stop
    # sudo apt-get purge mongodb-org*

    # 数据库和日志文件的路径取决于/etc/mongod.conf文件中的配置
    # sudo rm -r /var/log/mongodb   #移除日志文件
    # sudo rm -r /var/lib/mongodb   #移除数据库
```