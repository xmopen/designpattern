package creational

import "testing"

func TestNewSingleton(t *testing.T) {
	s1 := NewSingleton()
	s2 := NewSingleton()
	t.Logf("s1 ptr:[%p] s2 ptr:[%p]", s1, s2)
}
