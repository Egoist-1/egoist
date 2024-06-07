```
k8s
kubectl get all
kubectl apply -f dev/serv文件名.yaml   			//应用
会出现一行 deployment.apps/weboos configured
kubectl get deployments
kubectl get pods
kubectl get services
kubectl get namespaces
kubectl logs 
			-f  有一条就会添加一条
kubectl exec -it podId -- /bin/bash  //进入到对应pod的容器内
kubectl delete 组件+名称 
```



```
k8s
pod
	一个pod 就是一个实例,一个pod可以跑很多容器
service
	将一组pod封装成一个服务,并且提供统一的入口访问这个服务,只有deployment是无法从外面访问的,需要将pod 封装为一个逻辑上的服务,即service
deployment
	管理pod的组件
PersistentVolume(持久化卷)pv
	:告诉pvc自己需要什么类型,如果没有匹配就会报错
pvc
	:设置自己都有什么类型的,供pv选择
ingress:
	代表路由规则,分发到不用的service
ingressController
	ingress 时你的配置.ingressController时执行这些配置的
	
```