package user

import (
	"fmt"
	"sync"
	"time"
)

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

func GetAllUsersdeatils() {
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
}
