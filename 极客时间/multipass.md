multipass
```
multipass launch --name --cpus --memory --disk
multipass list
multipass start
multipass stop
multipass shell  //进入虚拟机
multipass exec  	//执行命令
multipass help
multipass 创建的虚拟机默认不允许ssh远程登录,如果想使用ssh登录
multipass shell  		
sudo apt install net-tools		
vi /etc/ssh/sshd_config
passwd
ssh-keygen
ssh-copy-id ubuntu@ip
alias k3s='ssh ubuntu@ip'
```
