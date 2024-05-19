package sync

import (
	"sync"
	"testing"
)

//测试三种形况
//读多写少		读写锁性能好
//读少写多		差不多,
//平均			读写锁好
func benchmark(b *testing.B,rw RW,read,write int){
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < read*100; i++ {
			wg.Add(1)
			go func ()  {
				defer wg.Done()
				rw.Read()
			}()
		}

		for i := 0; i < write * 10; i++ {
			wg.Add(1)
			go func ()  {
			    defer wg.Done()
				rw.Write()
			}()
		}
		wg.Wait()
	}	
}

func BenchmarkReadMore(b *testing.B){benchmark(b,&Lock{},9,1)}
func BenchmarkReadMoreRW(b *testing.B){benchmark(b,&RWLock{},9,1)}
func BenchmarkWriteMore(b *testing.B){benchmark(b,&Lock{},1,9)}
func BenchmarkWriteMoreRW(b *testing.B){benchmark(b,&RWLock{},1,9)}
func BenchmarkEqual(b *testing.B){benchmark(b,&Lock{},5,5)}
func BenchmarkEqualRW(b *testing.B){benchmark(b,&RWLock{},5,5)}