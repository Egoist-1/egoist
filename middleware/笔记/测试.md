```
wrk
wsl wrk -t1 -d1s -c2 -ssignup.luahttp://localhost:8080/users/signup
-t 线程数
-d 时间 1s 1m
-c 并发数
-s 测试脚本
```


```
测试:
mock
sqlmock go get github.com/DATA-DOG/go-sqlmock
```
