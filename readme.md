### 外卖平台

### 外卖平台
    端： 客户端
        商家后台
        配送端
        集团后台

#### 客户端 
* 简要包括：分类 
    * 包子粥铺
    * 能量西餐
    * 家常菜
    * 快食检餐
    * 暖胃粉面
    * 冒菜麻辣烫
    * 饮品 甜点
* 店铺展示：
    * 菜单 - 菜品展示 - 加入购物车 - 下单
    * 评价 - 互动功能 (评价与回复)
    * 商家 - 地址详情 以及定位 + 食品安全档案(资质展示) + 服务时间 + 活动展示
    * 排序：销量、速度、点评分高. 

#### 商家后台
    * 上架 下架
    * 活动设置
    * 交易信息
    * 接单通知
    * 门店信息
    * 客情分析
    * 评价回复

### 配送端
    * 接单通知
    * 配送路线指示
    * 薪资结算

### 活动设置
    * 平台的活动 - 可以随时关闭
    * 商家的活动：平台可以充钱进商家后台，以抵扣商家平台的抽佣，不抵扣配送佣金
    * 客户：平台充钱进商家店铺，作为活动经费，以30天为一个周期，冲入金额按照30天平均分配
        同时按照当日消费人数奖励，每天晚上24点准时结算。当日一人消费，则金额全返该消费者，
        当日多人消费，则平均分配到每个消费者中，若当日无人消费，则金额计入第二天累加作为第二天的活动经费。 


基于echo快速开发的模板框架 结构如下：

```text
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