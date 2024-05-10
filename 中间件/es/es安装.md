
```
ik 分词器
安装
进入docker容器内部
docker exec -it elasticsearch /bin/bash
在线下载并安装
./bin/elasticsearch-plugin  install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.12.1/elasticsearch-analysis-ik-7.12.1.zip
退出
exit
重启容器
docker restart elasticsearch
```


```
分词器作用:
    创建倒排索引时对文档的分词
    用户输入内容,的分词
ik分词模式
    ik_smary:只能切分
    ik_max_word:最细切片
ik分词器扩展词条:
    config目录 ikAnalyzer.cfg.xml 添加扩展词典和停用词典
```