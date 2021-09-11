# 阿里云SLS日志服务接入DEMO
演示如何使用github.com/toniz/otel把Go程序的日志上传到Aliyun的SLS服务。

## 目录说明
* server HTTP服务编写举例
* client 客户端Go程序编写举例
* pack1 pack2 演示HTTP服务调用外部包如何传递调用关系。

## 使用
* 参考aliyun的SLS服务手册创建project和service.
* 填写server和client的trace_config.json文件。
* 编译和运行

---
### 其它模块加入trace调用 
***redis模块***:
1. 头文件添加：
```
import redis "github.com/go-redis/redis/v8"
import "github.com/go-redis/redis/extra/redisotel/v8"
```
2. redis初始化:
```
rc = redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: "",
    DB:       0,
})
```
3. 给redis实例加上钩子
```
rc.AddHook(redisotel.NewTracingHook())
```
4. 正常使用redis各种指令就都会记录到trace.

----



## 结果演示：
调用结果在阿里云的SLS服务里面查看：
![result1](doc/调用结果1.png)  
--- 
![result2](doc/调用结果2.png)  
---  
![result3](doc/调用结果3.png)  
---


etc..

