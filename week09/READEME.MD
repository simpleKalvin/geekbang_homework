# 第九周

## fix length
> 核心在于客户端与服务端一起决定包体长度， 客户端根据约定最大包体长度传输，接收端根据包体长度解析读取


## fix delimiter
> 与 fix length 类似， 客户端与服务端一起决定一个定界符，客户端使用定界符作为多包分隔符，接收端根据分隔符切割数据

## LengthFieldBasedFrameDecoder
> 也是客户端与服务端使用共同的规则，不一样的是，约定的内容种类更多 字段偏移量， 字段长度， 调整长度，跳过长度


# GOIM DEBODE
> 根据ppt解析， 暂时没研究goim怎么启动，只能先根据PPT暂时解析包结构
![img.png](img1.png)