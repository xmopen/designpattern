package structural

import "testing"

func TestAdapter(t *testing.T) {
	adapter := &ConcreteAdapter{
		component: &ConcreteComponent{},
	}
	adapter.Do()
}
