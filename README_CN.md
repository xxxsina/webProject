# golang 学习

## 先 git 方法
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


#### 三者区别：

* go run 和 go build 后面是直接加xxx.go ; 从文件编译

* 而go install 后面是直接加xxx（xxx为目录名） ； 从文件夹编译


## 驱动支持

## LICENSE

BSD License
[https://www.cnblogs.com/Paul-watermelon/articles/10842752.html](https://www.cnblogs.com/Paul-watermelon/articles/10842752.html)
