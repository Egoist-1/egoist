package process

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func Test_Create_Process(t *testing.T) {
	//os 包的startProcess
	//os 包的 findProcess
	//返回$Path 中的位置
	ex, err := exec.LookPath("echo")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(ex)
	// 得到 *Cmd 实例后，接下来一般有两种写法：
	// 调用 Start()，接着调用 Wait()，然后会阻塞直到命令执行完成；
	// 调用 Run()，它内部会先调用 Start()，接着调用 Wait()；
	cmd := exec.Command(ex,"hello world")
	err = cmd.Run()
	fmt.Println(err)
}

//进程的属性和控制及
func Test_Process_Arrtibutes_Ctl(t *testing.T)  {
	pid := os.Getpid()
	fmt.Println(pid)
	//获取 实际用户id和组id
	uid :=os.Getuid()
	gid := os.Getgid()
	fmt.Println(uid,gid)
	//修改进程的当前目录
	err := os.Chdir("/home/chelly")
	fmt.Println(err)
	//进程的当前工作目录
	wd,_  := os.Getwd()
	fmt.Println(wd)
	//进程环境列表  子进程会继承父进程的环境 创建后，父子进程的环境相互独立，互不影响。
	envs := os.Environ()
	fmt.Println(envs)
	//获取指定的环境变量
	env := os.Getenv("GOPATH")
	env2,ok :=os.LookupEnv("GOPATH")
	fmt.Println(env)
	fmt.Println(env2,ok)
	//设置环境变量
	// os.Setenv()
	//Unsetenv 删除名为 key 的环境变量。
	//Clearenv 删除所有环境变量。
}
