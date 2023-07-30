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
