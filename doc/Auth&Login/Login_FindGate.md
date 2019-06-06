**找出有效GateIP的设计**

通常情况下，找出方式应该是找出CCU最少的那个Gate给玩家使用。
在下线某服务器的情况下，可以使得指定Gate上的用户越来越少，实现Gate服务器下线的平滑运维。


Auth&Login服务器在找出数据库中的gate，在返回给客户端前，

1. 如果当前玩家已经有登录信息存在，则尝试通知原来登录过的Gate服务器，让上一个用户下线。
1. 会尝试进行RPC.call通知相关Gate服务器有关玩家的loginToken和userid等相关信息。
1. 发送相关gate 的IP地址和Port

如果rpc调用失败，则删除数据库中相关数据，然后重新获取，直到找到有效Gate。
如果最后数据库数据中没有可用Gate数据，则给玩家返回错误信息。