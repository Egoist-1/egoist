.PHONY: docker
network:
# 提供了网络隔离,同一网络内的容器可以相互通信,
# 而不用网络的容器则不能直接通信,除非通过端口映射或其他网络配置
	@docker network ls
	@docker network create egoist
	@docker network rm egoist
up:
	@docker-compose -f mysql.yml 		up 		-d
	@docker-compose -f redis.yml 		up 		-d
	@docker-compose -f mongodb.yml 		up 		-d
	@docker-compose -f etcd.yml 		up 		-d
	@docker-compose -f es.yml 			up 		-d
	@docker-compose -f portainer.yml 	up 		-d
	@docker-compose -f kafka.yml		up		-d
down:
	@docker-compose -f mysql.yml 		down
	@docker-compose -f redis.yml		down
	@docker-compose -f mongodb.yml 		down
	@docker-compose -f etcd.yml 		down
	@docker-compose -f es.yml 			down
	@docker-compose -f portainer.yml 	down
	@docker-compose -f kafka.yml 	down