local key = KEYS[1]
local code = ARGV[1]

local r_code = redis.call("get",key)
local cnt = tonumber(redis.call("get",key..":cnt"))
if  cnt <= 0  then
    --用户一直输入错误,请重试
    return -1
elseif code == r_code then
    redis.call("set", cnt, -1)
    return 1
else
    --输入错误 -1
    redis.call("decr", cnt)
        return 0
end
