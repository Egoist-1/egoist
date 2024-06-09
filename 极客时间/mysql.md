```
upsert
insert on conflict() do nothing
insert on conflict() do updates ccc where
mysql只有一种写法 insert xxx on duplicate key updat xxx
```

```
索引
最朴素的方法 WHERE
怎么设计联合索引:where查询条件,最左原则,排序 范围 区分度esr

在设计order by语句的时候,要注意让order by命中索引
查询时只会用一个索引,那么这个索引最好包含order的列
如:where id  order by ctime utime  >==  联合 index id ct ut
```

```
事务
事务传播机制 依赖于threadlocal机制线程本地存储 ,go事务无法绑定与goroutinu,
goroutinu本身是不可操作的实体


事务每提交一次,数据库层面上都会执行一些操作:会根据刷日志/持久化配置,redolog undolog,binlog,把内存的数据同步到磁盘上直接刷新掉,所以批量处理只开一次处理业务提示性能
```





