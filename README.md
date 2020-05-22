# gRPC_demo
go语言学习grpc笔记


### 这是一个完整的grpc实例


服务端和客户端代码在example里
/impl 为接口实现代码
helloserver.exe文件和helloClient.exe文件可直接运行
示例如下：

服务端：
```
  D:\code\go\src\gRPC_demo>helloserver.exe
  2020/05/21 23:51:29 listen: 50001
```

客户端
```
D:\code\go\src\gRPC_demo>helloClient.exe
2020/05/21 23:52:48 connect success!
请输入名字：
China
请输入消息:
Good Good study!
Server Response: name: China  massage: Good Good study!
请输入名字：

```
