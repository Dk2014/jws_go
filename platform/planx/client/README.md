# Client 设计

处理客户端和服务器的数据传输

* 传输
  * 安全: 加密
  * 效率: 压缩
  * 完整性: 长度+CRC
* 有效性
  * 心跳
  * 初次连接客户端握手[^1]
* 统计
  * 收发数据量
  * 收发数据次数
  * 错误次数


[^1]: 客户端应该在初次连接成功后立即发出有效数据包进行握手（心跳？），否则应该视为无效，应该Close.

## Stories

0.0.1

* 完整性:长度校验，处理粘包
* 初次连接客户端握手
* 心跳维护

0.0.2
unknown

1.0.0
+CRC
+加密
+压缩
+统计

2.0.0
握手协议
断线重连：https://github.com/xjdrew/goscon
扩展Packet包协议

实现参考

* /MyProjects/golang/opensource/nsq/nsqd/client_v2.go


## Packet 设计

## 握手协议

客户端的心跳间隔等基本参数信息是否应该初次连接建立后由服务器来发送？握手协议？

通常情况下如果客户端升级TCP协议，通常是因为服务器升级协议。
如果客户端是旧协议，发送给新的服务器端，服务器端需要根据协议进行降级。

例如：

* 会不会出现有两种客户端出现，一种支持加密和压缩，另外一种不支持加密和压缩?
iOS和Android高端机支持，而Android低端机不支持。但是我们不希望给Android高低端机分服。
* 会不会同一个客户端有些数据需要压缩和加密，有些不需要？[^2]

[^2]: 这种情况通常应该做在Packet包基础协议里面，而不是握手协议里面



协议中可能包括：

* 服务器版本号
* 是否支持压缩
* 加密
* 心跳间隔要求

## Endian

C# SDK 默认用Little Endian,所以这里所有数据传输使用Little Endian.

## 心跳

心跳的定义：服务器应该定期收到Client的数据包或者专用心跳包（PING client -> server）.
服务器收到任何来自客户端的数据都认为客户端是alive的。

这种模式的心跳定义可以取代初次连接客户端握手。
这个心跳间隔直接使用SetReadDeadline来实现。
任何Read Timeout都视为玩家链接失去，服务器将会主动close connection.

客户端PONG：所有来自服务器的信息，都可以认为是PONG。都可以重置客户端心跳。客户端心跳只要不超时，则认为服务器是alive的，不会主动断开。


## 恶意发包撑爆服务器内存的攻击
`在完成整个需求之前先不考虑 1 Packet <--> 1 Message`

所以每个Packet都应该在合理的大小内？如果len强制限定再64K, 就是一种很好的保护。

但是游戏客户端和服务器的交互往往不能限定在这么小的数据上, 比如游戏启动过程中玩家客户端尝试获取玩家存档，如果一个请求发到服务器，服务器需要发给客户端的数据有可能达到500K~1M。

## 大数据包分拆
`N Packets <--> 1 Message`

无论你设定的数据Packet多大，很有可能需要设计这个功能。
需要考虑Packet Order, Packet Group


## Pool Packet
8K
16K
64K
128K
256K
512K
