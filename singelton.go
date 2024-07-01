package main

import "sync"

type Singleton struct {
	Value int
}

var (
	once     sync.Once
	instance *Singleton
)

func GetInstance(i int) *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Value: i,
		}
	})

	return instance
}
