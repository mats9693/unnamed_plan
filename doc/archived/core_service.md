# 核心服务

集中、统一管理微服务架构中，各个业务服务的实例，当前版本包含以下内容：

1. 配置中心：管理配置
2. 服务注册中心：管理每个业务服务实例的状态，以及请求分发

为什么将两个功能模块合并成一个服务呢？

1. 两个功能模块的作用都是**集中、统一管理某事物**，一个是配置，一个是业务服务实例的状态
2. 其他服务均依赖于这两个功能模块，如果能够合并成一个服务，则只需要处理一个服务的高可用
3. 两个模块间有联系（rc从cc读取配置，cc能够获取rc的地址以提供给业务服务），合并可以打破两个模块构成的循环依赖

## 配置中心(config center, cc)

### 功能

基础功能：

1. cc可以从数据库获取项目的全部配置
2. 所有服务都可以从cc获取自身的配置

高级功能：

1. 相同服务，支持不同等级(level)的配置，例如开发、生产等

### 设计

配置结构：

1. 服务配置：一个服务的完整配置，由**配置等级**和**一组配置项**组成，举个例子：**用户服务的开发模式配置**
2. 配置项：一个模块的配置，由**配置项标识**和**一系列键值对**组成，举个例子：**数据库配置项**
    1. 配置项标识：包含配置项ID、配置项名称
    2. 配置项的键值对：配置值，例如数据库配置项的`"DBMS": "postgresql"`
3. 配置项可以复用，即不同服务或不同等级的相同服务，可以包含同一个配置项

cc在启动时加载全部配置，并计算出支持哪些服务的哪些配置等级，当前版本支持**默认**、**开发**、**生产**、**测试**配置等级  
其他服务携带**服务ID**、**配置等级**访问cc，如果cc支持该服务，但不支持该配置等级，则返回**默认**等级的配置

### 问题

1. 当前通过一个**服务公共配置文件**提供cc的地址，其实可以省略这个配置文件，在**服务注册中心-嵌入模块**
   (Registration Center Embedded, RCE)中，硬编码cc的域名，然后通过部署手段将域名与cc联系起来（关键词：域名解析、反向代理）

### 优化

1. 编写配置中心管理界面：为了方便修改配置
2. 配置更新同步：配置有更新，如何同步到目标服务呢？
3. 配置导出到文件：为了方便管理业务服务测试配置
4. 配置版本管理：拟将配置版本加入管理，例如**用户服务1.0版本配置**、**用户服务2.0版本配置**

## 服务注册中心(registration center, rc)

### 功能

1. 业务服务向rc注册
2. 其他服务从rc获取指定业务服务可用实例的地址列表
3. 维护实例状态（是否可用，rcc/rce）
4. 其他服务调用业务服务

### 设计

rc包括**核心模块**和**嵌入模块**

1. 核心模块(registration center core, RCC)：服务端模块，负责响应业务服务注册、对其他服务提供服务实例查询、维护实例状态
2. 嵌入模块(registration center embedded, RCE)：嵌入到其他服务中，负责向rcc注册、调用其他业务服务
    1. 其他服务：业务服务（用到注册功能，可能用到调用功能，例如用户服务、云文件服务）和需要调用业务服务的服务（用到调用功能，例如网关）

> 在代码中，我们使用`rce`表示**服务注册中心-嵌入模块**，直接使用`rc`表示**服务注册中心-核心模块**

rcc启动，开始维护实例状态  
当有业务服务启动时，其内嵌的rce模块向rcc模块注册该服务  
当一个服务需要调用某个业务服务时，先查看自身内嵌的rce模块是否持有目标服务的实例地址，如果没有，则从rcc模块获取；如果有，选择一个地址发起调用

以下过程中，可能出现实例不可用的情况（grpc带给go的error不空），当前版本的处理规则是：直接删除不可用实例

1. rcc模块维护实例状态
2. rce模块调用业务服务

### 问题

1. rcc重启可用：现在的rcc重启，已经注册的实例信息将被清空。考虑rce重复注册

### 优化

1. rce，选择实例地址的方式，考虑升级为**根据响应时间计算权重**
2. rce，维护持有的业务服务实例地址列表的方式，考虑升级为
   **出现网络错误时，将该实例降级，如果所有地址都降级了，则重新向核心模块请求**
3. rce，拟于rpc请求因网络原因失败时，增加**切换地址**、**重试**功能
4. rcc，维护业务服务实例地址列表的规则，考虑升级为
   **出现网络错误时，将该实例降级，收到rce请求时，返回正常等级实例，如果没有，则返回低级实例**
5. rcc，获取到的目标服务地址，如果与当前服务处于同一局域网内，则使用内网ip访问（没想好怎么判断内网，ip前16个bits吗）
