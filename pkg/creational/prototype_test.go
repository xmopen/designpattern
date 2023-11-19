package creational

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	// 创建一个具体原型
	prototype := &ConcretePrototype{
		Value: 100,
	}
	// 克隆原型
	clone := prototype.Clone()
	prototype.Value = 200
	t.Logf("prototype: [%+v]", prototype.GetValue())
	t.Logf("clone:[%+v]", clone.GetValue())
}
