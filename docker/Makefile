.PHONY: docker
network:
# 提供了网络隔离,同一网络内的容器可以相互通信,
# 而不用网络的容器则不能直接通信,除非通过端口映射或其他网络配置
	@docker network ls
	@docker network create egoist
	@docker network rm egoist