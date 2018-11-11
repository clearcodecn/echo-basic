### echo basic 

基于echo快速开发的模板框架 结构如下：
```bash
├── app 
│   ├── controller    # 控制器
│   ├── httpio       
│   │   ├── in        # 请求结构体定义  
│   │   └── out       # 响应结构体定义  
│   ├── middleware    # 自定义中间件， 错误处理程序 
│   ├── model         # 持久层   
│   ├── repo          # db层封装
│   └── router        # 路由定义
├── cmd
│   └── server        # 启动程序
├── config
├── module            # module 模块化  
│   ├── webserver     # http服务
│   └── wsserver      # websocket服务  
└── pkg               # 自定义的扩展包 
    ├── atomicbool    # 原子bool
    ├── modules       # 模块包
    └── validator     # 验证包

```

#### 依赖

* go 必须大于等于 `1.11`

```bash

github.com/spf13/viper  配置管理
github.com/labstack/echo http服务器
github.com/urfave/cli 命令行工具
github.com/asaskevich/govalidator 验证库
github.com/gorilla/websocket websocket连接库
github.com/jinzhu/gorm 数据库orm
github.com/gomodule/redigo/redis redis


```

#### 使用代理下载`golang.org/x`：
```bash
export GOPROXY=https://goproxy.io
```

#### 运行

```bash
# export GOPROXY=https://goproxy.io  可以代理下载 golang.org/x 包
go run cmd/server/server.go 
# or 
go run cmd/server/server.go -c otherCfg.yaml server
# or 
go run cmd/server/server.go -c otherCfg.json server

```