package time

import (
	"fmt"
	"testing"
	"time"
)
func Test_Time(t *testing.T)  {
	now := time.Now()
	//时间减法
	result := now.Add(time.Hour * -5)
	fmt.Println(result)
	n2 := now.Sub(result)
	fmt.Println(n2)
	//解析  出来的是utc时间 会与 time.now() 的local 不符预期
	t2,_:= time.Parse("2006-01-02 15-04-05", "2024-05-14 18-08-58")
	fmt.Println(t2)
	//解析 应使用
	t2,_= time.ParseInLocation("2006-01-02 15-04-05", "2024-05-14 18-08-58", time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Format("2006-01-02 : 15"))
	t3, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:00:00"), time.Local)
	fmt.Println(t3)
	//取整 round truncate
}


func Timer(t *testing.T) {
	
}
