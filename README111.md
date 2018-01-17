# mysql订阅服务

## 组成
- mysql监控服务 
- 监控管理系统
- 转发服务
- A elasticsearch 检索服务（未完成）
- B mysql 触发器服务 （未开始做）
- C http 接口服务

### mysql监控服务
```
- 主要的 第三方模块 https://github.com/zalora/binlog-parser.git
1. 监控mysql的binlog日志
2. 对mysql的 dml dll 触发器 等操作 解析
3. 将解析完的数据存放到redis(以数据库名为key)/将解析完的数据存放到kafka(以数据库名为topic)
说明: 3 初级版可以使redis， 如果数据量大使用kafka
```
### 监控管理系统
- web服务
- 使用http请求 跟 监控服务，转发服务 通信
- 管理系统 通过 http请求 在监控服务上设置监控哪些库 哪些表
- 管理系统 通过 http请求 向转发服务发送 更新监控数据 的命令 

### 转发服务
```
1. 读取redis/kafka中的数据 分发给 ABC 3个服务
2. AB 是基础平台使用 可以提高接口查询速度 
3. AB完成之后，基本mysql查询的接口将全部废弃
4. C 为各个项目提供的接口服务 便于数据同步
```

### ABC
```
1. ABC依赖于转发服务
2. BC的数据源就是转发服务
3. A的数据源是mysql的旧数据和转发服务提供的更新数据
```
