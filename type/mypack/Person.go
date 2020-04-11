package mypack

import "fmt"

type Person struct {
}

//类型别名
type Student = Person

func (p *Person) Run() {
	fmt.Println("run...")
}

func (p *Student) Study() {
	fmt.Println("study...")
}
