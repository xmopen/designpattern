package creational

// 适用场景：当需要在多个场景中进行同一类型数据的共享时可以通过原型模式来进行共享

// Prototype 原型接口
type Prototype interface {
	Clone() Prototype
	GetValue() int
}

// ConcretePrototype 具体原型
type ConcretePrototype struct {
	Value int
}

func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		Value: c.Value,
	}
}

func (c *ConcretePrototype) GetValue() int {
	return c.Value
}
