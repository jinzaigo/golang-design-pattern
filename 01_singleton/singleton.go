package main

import "fmt"

type singelton struct{}

var instance *singelton = new(singelton)

func GetInstance() *singelton {
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例对象的某方法-饿汉式")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
