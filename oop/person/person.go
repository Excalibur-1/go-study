//封装：把抽象出的字段和对字段的操作封装在一起,数据被保护在内部,
//程序的其它包只有通过被授权的操作(方法),才能对字段进行修改，其作用有：
//	隐藏实现细节
//	可以对数据进行验证，保证安全合理
//Golang对面向对象做了极大简化，并不强调封装特性，下列示例进行模拟实现：
package person

import "fmt"

type Person struct {
	Name string
	age  int //年龄是隐私，不允许其他包访问
}

//工厂函数（类似构造函数）
func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

func (p *Person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄不合法")
	}
}

func (p *Person) GetAge() int {
	return p.age
}
