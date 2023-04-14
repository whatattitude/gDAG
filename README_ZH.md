[英文](README.md) | 中文

# gDAG
gDAG是一个用于调度和监控工作流的平台。gDAG是一个任务调度工具，使用有向无环图(DAG)管理任务流。任务调度可以在不了解业务数据内容的情况下，通过设置任务依赖关系实现。


## 目录结构
- control.sh 启动控制文件
- app 调度器，worker等app入口
- config 测试配置目录
- bin  二进制产出
- docs  项目说明文档
- lib 基础库，相互之间不依赖，不依赖外层
- service 逻辑服务层

  
## control.sh 功能说明
1.构建app, 不带build参数默认构建app目录下所有应用，带参数构建指定的应用，构建产出到bin目录下
> `#sh control.sh build dashboard`

2.清理bin下所有产物
> `#sh control.sh clean`

3.在8001端口启动本地doc文档
> `#sh control.sh doc`
> `http://localhost:8001/pkg/gDAG/`