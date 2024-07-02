package singelton

import "sync"

type singleton struct {
	Value int
}

var (
	once     sync.Once
	instance *singleton
)

func GetInstance(i int) *singleton {
	once.Do(func() {
		instance = &singleton{
			Value: i,
		}
	})

	return instance
}
