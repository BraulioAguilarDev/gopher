package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *DriverPg
)

type DriverPg struct {
	conn string
}

func Connect() *DriverPg {
	once.Do(func() {
		instance = &DriverPg{conn: "driverConnectPg"}
	})
	return instance
}

func main() {
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*Connect())
	}()

	for i := 0; i < 100; i++ {
		go func(i int) {
			time.Sleep(time.Millisecond * 600)
			fmt.Println(i, "=", Connect().conn)
		}(i)
	}
	fmt.Scanln()
}
