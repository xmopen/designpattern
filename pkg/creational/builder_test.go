package creational

import "testing"

func TestBuilder(t *testing.T) {
	builder := &ConcreteBuilder{}
	t.Logf("result:[%+v]", builder.Build().String())
}
