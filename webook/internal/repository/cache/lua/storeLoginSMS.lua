local key = KEYS[1]
local val = ARGV[1]
local keyCnt = key..":cnt"
--ttl -1 存在但没有过期时间 -2 kye不存在
local ttl = tonumber(redis.call("ttl",key))
if ttl == -1 then
    --存在但没有过期时间
    return -1
elseif ttl == -2 or ttl <= 560 then
    --正常
    redis.call("set",key,val)
    redis.call("expire",key,600)
    redis.call("set",keyCnt,3)
    redis.call("expire",keyCnt,600)
    return 1
else
    --发送太频繁
    return 0
end