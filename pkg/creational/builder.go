package creational

import "encoding/json"

// 构造者模式
// 适和使用场景
// 1、当需要创建复杂对象时
// 2、当需要实现不同类型的产品时
// 3、解耦，对象的创建和对象的实现分离

// 优点：
// 1、解耦：对象的创建和表示进行分离，使的同样的构造过程可以创建不同的对象的表示
// 2、扩展：可以更加方便的添加新的构造过程，而无需修改原有代码
// 3、灵活：可以根据需要选择不同的构造过程来创建不同的表示

// 缺点：
// 1、代码量增加
// 2、代码可读性降低
// 3、构造的对象和构造的过程的匹配，需要明确知道自己需要哪种构造器来创建自己所需要的对象

// Product  interface
type Product interface {
	String() string
}

// ProductImpl product impl
type ProductImpl struct {
	Name string
}

// String marshal ProductImpl to json string
func (p *ProductImpl) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}

// Builder builder itnerface
type Builder interface {
	Build()
}

// ConcreteBuilder builder
type ConcreteBuilder struct {
}

// Build product instance
func (c *ConcreteBuilder) Build() Product {
	return &ProductImpl{
		Name: "product",
	}
}
