# readme
# 文件夹说明
+ `go.mod`
+ `rpc` rpc服务
+ `api` api http调用
+ `cronjob` task任务
+ `script` 脚本
+ `rmp` 队列
+ `model` 数据库

文件夹tree说明
```
xxx-service
|____go.mod
|____README.md
|____rpc
| |____xxxxx
|____api
| |____xxxxx
|model
| |____xxxxx
|cronjob
| |____xxxxx
|script
| |____xxxxx
|rmp
| |____xxxxx
```

# idea配置
## 自动引用mac系统的mod库
1. 点击菜单栏 Goland
2. Preferences
3. Go -> Go Modules
4. 打钩✅Enable Go modules integration
5. 配置GOPROXY 选择 direct保存

## 创建mod
```
 go mod init github.com/pz2147/p-rpc-1
```

# RPC 规范
1. 生成模板
```
goctl rpc template -o=pRpc1.proto  
```
3. 编辑`pRpc1.proto`定义
   1. package 规范
    ```
    # 使用全小写
    package pRpc1
    ```
   2. service 规范
    ```
    service PRpc1 {}
    ```
4. 生成rpc文件夹路径
```
_goctl rpc proto -src guildService.proto -dir . -style goZero_
```
也可以根目录下，使用makefile命令
```
make goctl-rpc
```

# 数据库规范
1. 根目录`cd model`
```
goctl  model mysql datasource -url="account:password(db_ip:3306)/db_name" -table="table_name"  -dir="." -style goZero 
```
+ 生成对应的table_name的数据库文件.go
+ 切记代码风格`-style goZero`

----

# 文件夹说明
无论是api/rpc生成都有对应的internal文件夹
## etc配置
1. `./etc/xxxService.yaml` 配置数据库,redis,mq位置信息
## 内部的文件夹管理
1. `./internal/config/config.go` 配置数据库，redis，其他外部组件信息
2. `./internal/svc/serviceContext.go` 配置数据库关系，外部rpc调用关系
3. `./internal/logic` 编写代码逻辑


# idea配置debug
1. mac的idea
2. 点击右上角的add configuration打开界面
3. 点击左上角"+"
4. 选择go build
 
## 以api为例（rpc一样的）  
5. package path 相对路径
   eg:
   package path: code.wecochat.com/mesa/guild-service/api
6. working directory 项目的绝对路径
   eg:
   working directory: 绝对路径../guild-service/api
7. 点击 play 或者 debug,



# swagger
goctl api plugin -plugin goctl-swagger="swagger -filename guildService.json" -api guildService.api -dir .

scp ./guild.json root@47.119.147.135:/data/script/config/swagger/


### 用户信息
成长体系：

公会用户列表需求：
1、财富等级
2、活跃等级
3、魅力值
4、用ID  uid
5、昵称  nickname
6、头像  avatar_url
7、性别  gender

----------
8、钻石数
---分页---
9、page_size
10、page
11、total

现在排序需求： 月度钻石最多 ，从加入公会开始，获得的所有钻石



加多一个，申请的成员列表


### etcd
./etcdctl get "" --from-key

### 获得ES所有Index
http://120.79.67.15:9200/_cat/indices?v&pretty