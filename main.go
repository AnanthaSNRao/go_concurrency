package main

import (
	"fmt"
	"sync"

	ms "github.com/go-concurrency/messageServer"
	singelton "github.com/go-concurrency/singelton"
	tsmap "github.com/go-concurrency/threadsafemap"
	user "github.com/go-concurrency/user"
)

func main() {
	user.GetAllUsersdeatils()

	safeMap := tsmap.NewSafeMap()

	// Store a value
	safeMap.Set("key1", "value1")

	// Load a value
	if value, ok := safeMap.Get("key1"); ok {
		fmt.Println("Loaded value:", value)
	}

	// Delete a value
	safeMap.Delete("key1")

	// Load a value after deletion
	if value, ok := safeMap.Get("key1"); !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println("Loaded value:", value)
	}

	g := singelton.GetInstance(9)

	fmt.Printf("fist instance of singelton: %v \n", g)

	f := singelton.GetInstance(10)

	fmt.Printf("Second instance of singelton: %v \n", f)
	msgch := make(chan ms.Message)

	s := ms.GetInstance(msgch)

	wg := &sync.WaitGroup{}

	go s.StartAndListen(wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go ms.SendMessageToServer(msgch, fmt.Sprint(i), fmt.Sprint(i*i))
	}
	wg.Wait()
	go func() {

		s.ShutDownServer()

	}()

}
