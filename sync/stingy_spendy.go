package main

import (
	"time"
)

var money = 100

// var lock = sync.Mutex{}

func stingy() {
	for i := 1; i <= 100; i++ {
		// lock.Lock()
		money += 10
		// lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}
func spendy() {
	for i := 1; i <= 100; i++ {
		// lock.Lock()
		money -= 10
		// lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	println(money)
}
