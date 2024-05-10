FROM node:14-alpine

#  source: 相对于Dockerfile  dest:相对于镜像的路径
COPY  index.js /index.js  

# 运行程序 1:可执行程序的名字 2:接受的参数 
# CMD [ "node" ,"/index.js"]  
# ENTRYPOINT [ "" ]
CMD node /index.js


# 构建
# docker build -t hello-docker:v0.0.1 .  

# 容器
# docker run name
# docker stop [options] container [container...] ]停止容器
# docker ps -a  列出所有容器包括停止的
# docker rm  {id}  移除容器

# 镜像
# docker rmi {name}  移除镜像
# docker pull 地址 下载镜像
