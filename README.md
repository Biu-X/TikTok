# 第六届字节青训营后端大项目-抖音

## 一、背景
1. **项目介绍**  
一句话介绍，实现极简版抖音。每一个应用都是从基础版本逐渐发展迭代过来的，希望同学们能通过实现一个极简版的抖音，来切实实践课程中学到的知识点，如 Go 语言编程，常用框架、数据库、对象存储等内容，同时对开发工作有更多的深入了解与认识，长远讲能对大家的个人技术成长或视野有启发。

## package 依赖关系
> 由于Golang不支持循环依赖，我们必须仔细决定包之间的依赖关系。这些包之间有一些级别。以下是理想的包依赖关系方向。  

`cmd` -> `routers` -> `services` -> `models` -> `modules`

从左到右，左边的包可以依赖右边的包，但是右边的包不能依赖左边的包。 同一级别的子包可以根据该级别的规则进行依赖。  


## 作者
- [@Tohrusky](https://github.com/Tohrusky)
- [@Zaire404](https://github.com/Zaire404)
- [@resortHe](https://github.com/resortHe)
- [@KelinGoon](https://github.com/KelinGoon)
- [@1055373165](https://github.com/1055373165)
- [@dwe321](https://github.com/dwe321)
- [@Isaac03914](https://github.com/Isaac03914)
- [@hiifong](https://github.com/hiifong)

```text
tiktok
├── LICENSE
├── Makefile # make 命令默认显示帮助信息, 开发期建议使用make watch可以热加载代码, 注意 make 命令不推荐 Windows 用户使用
├── README.md
├── TikTok.go
├── apk
├── cmd
│     ├── cmd.go # 命令行入口
│     ├── gen.go # gen 子命令,用于从 MySQL 生成 gorm 相关的代码
│     └── web.go # server 子命令,用于启动后端接口服务器
├── conf
│     └── tiktok.yml # tiktok 的配置文件,例如 mysql 的配置
├── dal
│     └── query # gen 子命令自动生成的代码,用于 CRUD
│         ├── gen.go
│         ├── gen_test.db
│         ├── gen_test.go
│         ├── users.gen.go
│         └── users.gen_test.go
├── docs
├── go.mod
├── go.sum
├── models # gen 子命令从 MySQL 生成的 model
│     └── users.gen.go
├── modules # tiktok 的各种模块
│     ├── config
│     │     ├── config.go # 初始化配置模块,后续可以通过`config.Get(key string)/config.GetString()`获取配置
│     │     └── config_test.go
│     ├── context # 上下文,暂未设计好
│     │     ├── context.go
│     │     └── context_test.go
│     ├── db
│     │     └── db.go # 用于初始化数据库及数据库实例
│     ├── ffmpeg
│     │     └── ffmpeg.go # 获取视频截图用作视频封面
│     ├── gen
│     │     └── gen.go # gen 模块,用于配置g en 自动生成代码
│     ├── govatar
│     │     ├── README.md
│     │     ├── govatar.go # govatar 通过 email 生成头像
│     │     ├── govatar_test.go
│     │     └── hiifong.png
│     ├── middleware # 中间件模块
│     │     ├── jwt
│     │     │     ├── jwt.go
│     │     │     └── jwt_test.go
│     │     └── middleware.go
│     └── redis # Redis 模块, 初始化 Redis 及提供 Redis 实例
│           ├── redis.go
│           └── redis_test.go
├── routers
│     ├── api
│     │     └── v1
│     │         └── api.go # 接口模块,相当于 controller
│     └── init.go # 初始化接口
├── scripts
│     ├── MySQL.sh
│     └── tiktok.service
└── services # 业务模块,例如登录,注册等业务模块
    ├── auth
    │     └── auth.go
    └── user
        └── user.go
```
