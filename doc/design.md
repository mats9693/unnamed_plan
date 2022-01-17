# 设计

## API 网关

只接受http请求

### http请求响应过程

1. 判断**请求来源**，若请求来源无效（未注册），则拒绝该请求
    1. 请求来源：http请求头```Unnamed-Plan-Source```字段的值
    2. 启动时注册有效的来源列表
2. 获取对应请求的处理函数
    1. 所有处理函数在程序启动时注册，此处根据URI获取
3. 多处登录限制：验证（多处登录限制详见下文）
    1. 满足下列条件之一的，跳过验证：
        1. 未开启多登限制
        2. 请求来自**非限制来源**
        3. 针对该处理函数**不启用**多登限制（注册处理函数时设置）
    2. 执行验证并通过后，刷新用户在该来源的**上一次请求时间**
        1. 验证不通过的请求将被拒绝
4. 调用处理函数
5. 多处登录限制：刷新登录信息
    1. 满足下列条件之一的，不刷新登录信息：
        1. 未开启多登限制
        2. 请求来自**非限制来源**
        3. 该处理函数**未启用**刷新token功能（注册处理函数时设置）
        4. 处理函数执行失败
    2. 刷新登录信息会同时刷新```token```和**上一次请求时间**

### http res

```go 
type ResponseData struct {
	HasError bool   `json:"hasError"`
	Data     string `json:"data"`
}
```

```hasError```表示http请求是否正确执行并按照预期返回：

1. 请求正常执行：```data```是```json string```
2. 请求执行失败：```data```是```string```类型的错误信息

### 多处登录限制 multi-login limit

#### 目的

限制同一用户同时登录多个终端的情况

#### 策略

1. 同一用户，在同一个**限制登录**的来源上，**最多**只能登录一次
2. 当一名用户尝试在同一个来源的**不同终端**上登录时，先登录的终端将被收回登录权限
    1. 不同终端：例如两个浏览器标签页

#### 实现

1. 前端将```token```保存到```sessionStorage```
2. 前端通过自定义的http请求头携带身份验证参数
    1. 来源：```Unnamed-Plan-Source```
    2. 用户ID：```Unnamed-Plan-User```
    3. token：```Unnamed-Plan-Token```
3. 满足以下全部条件的，视为验证通过：
    1. 解析出有效的```user id```与```token```
    2. 程序内保存有该```user id```的登录信息
    3. 该用户在该来源的**上一次请求时间**距**此刻**不超过一定值（可配置）

#### 问题

1. 考虑引入缓存层保存用户登录信息
2. 后端识别**请求头**中的关键参数，然后把关键参数写在**请求体**里返回，  
   这种做法能够最大程度兼容```go http server```和```vue axios interceptor```的原生写法  
   简单来说就是容易实现，后续考虑统一
    1. 考虑使用服务端设置cookie的方式，涉及跨域问题

## 配置中心

> 第一版不引入服务注册中心，跨服务调用方式保持现状

目的：解决多个服务配置分散、不容易修改和管理的问题

设计：（下面以**用户服务**代表实现某种业务的服务）

1. 用户服务持有**用户服务ID**与**配置中心地址**
    1. 用户服务ID：固定在代码中；当有多个用户服务实例时，它们的服务ID相同、实例ID由服务注册中心分配
    2. 配置中心地址：通过配置文件给出，仅包含一条配置项，所有服务实例均可读相同的配置文件
2. 用户服务访问配置中心
    1. 用户服务连接配置中心，根据**服务ID**获取**服务配置**（包括服务注册中心地址）
    2. 用户服务根据获得的配置启动，启动成功，向**服务注册中心**注册自身，服务注册中心返回**实例ID**并开始维护实例状态
3. 配置中心有调整配置的可视化界面，界面包含以下功能：
    1. 维护服务ID，例如**增加云文件服务ID**等操作
        1. 查询、新增、修改、删除服务ID
    2. 维护具体配置，例如**修改数据库配置连接地址**等操作
        1. 查询、新增、修改、删除配置项
    3. 维护服务ID对应配置，例如**为用户服务增加数据库配置**等操作
        1. 查询、新增、删除对应关系
    4. 维护配置版本，例如**读取上个版本的配置**等操作
        1. 生成新版本配置、查询历史配置、应用其他版本配置
    5. 测试读取配置，例如**测试用户服务能读取到的配置是否合理**等操作
        1. 再想一想做不做
4. 配置中心有三张表：
    1. 服务ID表：包含所有服务ID，以及每个服务ID对应的配置ID列表
    2. 配置详情表：包含所有配置ID，以及每个配置ID包含的配置详情
    3. 成版本的完整配置表：包含历史配置版本，以及每个版本的完整配置
5. 配置中心从数据库读取配置，启动时加载配置表中标记为**当前使用**的版本的配置，应用其他版本配置时，修改中心缓存
    1. 思考：如何保证从配置中心获取的配置有效？

## 服务注册中心(draft)

目的：解决业务服务实例数量提升所带来的新问题，例如负载均衡、实例状态维护等

1. 服务注册中心实现所有业务服务的客户端代码与转发代码
    1. 调用链调整为：API网关/其他服务 -> 服务注册中心 -> 目标服务
    2. 服务注册中心应实现全部rpc方法的服务端与客户端，即，服务注册中心在收到请求时，仅做转发
    3. 是否要支持**查询服务注册中心当前支持的业务函数**呢？
    4. note：跨服务调用通过服务注册中心，是因为这样调用方就不需要知道目标服务地址了，负载均衡、目标服务状态等都不需要关心

## 性能监控(draft)

pprof

## 调用链追踪(draft)

日志收集，一个请求在各环节拥有同一个请求ID

执行业务的服务中增加日志

rpc方法拟使用自定义错误类型

## 熔断、限流、降级(draft)

熔断、降级拟合并到服务注册中心，限流拟合并到API网关