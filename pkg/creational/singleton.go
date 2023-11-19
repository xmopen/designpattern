package creational

import "sync"

var (
	singletonInstance *Singleton
	singletonOnce     sync.Once
)

// Singleton 单利模式
type Singleton struct {
	Value int
}

func NewSingleton() *Singleton {
	singletonOnce.Do(func() {
		if singletonInstance == nil {
			singletonInstance = &Singleton{
				Value: 1,
			}
		}
	})
	return singletonInstance
}
