package structural

import "fmt"

// 适配器模式：当接口不兼容时，通过适配器模式，将一个类的接口转换成客户希望的另一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。
// 优点：
// 1、将不同的类或者接口协同工作
// 2、增加代码的扩展性
// 缺点：
// 1、增加系统复杂性

// Component 被适配接口
type Component interface {
	Do()
}

// Adapter 适配接口
type Adapter interface {
	Do()
}

// ConcreteComponent  一个具体被适配的类
type ConcreteComponent struct {
}

func (c *ConcreteComponent) Do() {
	fmt.Println("ConcreteComponent")
}

// ConcreteAdapter  具体适配器
type ConcreteAdapter struct {
	component Component
}

func (c *ConcreteAdapter) Do() {
	c.component.Do()
}
