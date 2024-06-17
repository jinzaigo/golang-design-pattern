// 饿汉式：初始化时就创建实例，优点简单、线程安全，缺点是可能会浪费内存
package main

import (
	"fmt"
	"sync"
)

type singelton struct{}

var (
	instance *singelton
	once     sync.Once
)

func GetInstance() *singelton {
	if instance == nil {
		once.Do(func() {
			instance = new(singelton)
		})
	}
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例对象的某方法-懒汉式")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
