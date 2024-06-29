package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	wg1 := &sync.WaitGroup{}
	respch := make(chan string, 3)

	go getUserDetails(1, respch, wg1)
	wg1.Add(1)
	go getUserReccomendatations(1, respch, wg1)
	wg1.Add(1)
	go getUserPermissions(1, respch, wg1)
	wg1.Add(1)

	wg1.Wait()

	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))

	safeMap := NewSafeMap()

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

	s := Server{
		msgch: make(chan Message),
	}
	wg := &sync.WaitGroup{}

	go s.StartAndListen(wg)

	for i := 0; i < 10; i++ {
		go sendMessageToServer(s.msgch, fmt.Sprint(i), fmt.Sprint(i*i))
		wg.Add(1)
	}
	wg.Wait()
	go func() {

		shutDownServer(s.quitch)

	}()

}

func getUserDetails(userId int, respch chan string, wg *sync.WaitGroup) {

	time.Sleep(80 * time.Millisecond)

	respch <- fmt.Sprintf("user Deaitls: %v", userId)

	wg.Done()
}

func getUserReccomendatations(userId int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)

	respch <- fmt.Sprintf("user Reccomendatation: %v", userId)

	wg.Done()

}

func getUserPermissions(userId int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)

	respch <- fmt.Sprintf("user Permissions: %v", userId)

	wg.Done()
}
