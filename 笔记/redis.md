```
使用redis的zset命令实现滑动窗口
假设窗口时 1分钟 限流是n 使用时间戳作为score 
来了一个请求 此时时间戳是t1
看看zset score(t1-1m,t1)之内有多少请求
如果 <n 就放行
否则 拒绝
```


```
prometheus 监控 redis 的缓存命中率 redis.nil
```