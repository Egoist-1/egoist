```
为什么要使用长短token
效率和安全的折中
因为短token 频繁的被使用 容易泄露
所以要是短token过期了,直接生成一个新token就和刷新区别不大了
因此只有长token调用refresh时 才会生成新的短token
提问:长token泄露了
	回答:长短token同时刷新
		或者可以在长token里加入user-agent其他校验是否时本人
长短tokne同时刷新存在一个问题,调用refresh时超时了,这时老的refresh已经没用了所以用户只能重新登录
```
