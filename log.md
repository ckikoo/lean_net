# dialtimeout 普通的dial qubie
>  当出现不可以到达时
>  超时 io.timeout  

# server 没有及时accept 
> 阻塞处理，没异步--》 出现的问题，要么超时，要吃报错

> backlog 也就是半连接队列的问题
> window 200个大小
> linux 128 个附近

#  read 和write 的注意事项
## write 特点
* 写成功 err =nil  && wn == len(buf)
* 写阻塞
*  使用deadline 来控制超时时间

## read 特点
* conn 无数据, read 阻塞
* 数据过大,开缓存区
* 部分, 返回当前读取长度
* 关闭 eof   (全双工, 既可以发送也可以读取)
* 设置时间
* deadline 