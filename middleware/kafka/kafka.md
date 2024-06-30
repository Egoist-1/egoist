#监听test_topic,链接localhost:9094

kafka-console-consumer -topic=test_topic -brokers=localhost:9094

在kafka,生产者发送数据,一个很关键的参数acks
0:客户端发送一次,不需要服务端确认 ,tcp协议返回了 ack 就可以了
1:客户端发送,并且需要服务端写入到主分区  ,主分区确认写入就可以了
-1:客户端发送,并且需要服务端同步到所有ISR(In sync Replicas),所有的ISR都确认
从上到下,性能变差,可靠性上升.需要性能 0 ,需要消息不丢失 -1