package test

import (
	"sync"
)

//table test data
func Fib(n int) int {
	if n < 2 {
			return n
	}
	return Fib(n-1) + Fib(n-2)
}

//Parallel test data
var(
	data = make(map[string]string)
	locker sync.RWMutex
)
func WriteToMap(k,v string)  {
	locker.Lock()
	defer locker.Unlock()
	data[k] = v
}
func ReadFromMap(k string)string {
	locker.Lock()
	defer locker.Unlock()
	return data[k]
}

func Add(a int, b int) int {
    return a + b
}

func Mul(a int, b int) int {
    return a * b
}

// db.go
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}