### V1版本

```
v1版本主要做了服务器的链接,无太多功能
```

- 注意点

  - go run 之前先使用 `go mod init im-server` 其中 "im-server"是module的名称,自定义的

  - 调试时不能用 `go run main.go`,因为只编译了`main.go`文件,没有编译到`server.go`文件,会报错;推荐使用`go run .`用来调试

  - 由于没有前端链接,推荐使用 nc或者curl来使用

    ```
    nc先执行命令,安装nc命令,如果是其他系统请跳过或自行寻找相关教程
    apt-get update
    apt-get install netcat
    
    nc 127.0.0.1 8000
    或者
    curl 127.0.0.1:8000
    
    这时候执行go run . 的服务端会提示
    root@2b7fa705c039:/goCode/goChatDemo# go run .
    链接创建成功
    ```

    
