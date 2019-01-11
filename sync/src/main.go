package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	doInCaseSimpleChannel()
	fmt.Println()

	doInCaseMutex()
	fmt.Println()

	doInCaseOnce()
	fmt.Println()

	doInCaseWaitGroup()
	fmt.Println()

	doInCaseCond()
	fmt.Println()
}

func doInCaseSimpleChannel() {
	fmt.Println("doInCaseSimpleChannel start...")

	c := make(chan bool)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
			c <- true
		}()
	}
	for i := 0; i < 5; i++ {
		<-c
	}

	fmt.Println("finish!")
}

func doInCaseMutex() {
	fmt.Println("doInCaseMutex start...")

	c := make(chan bool)
	m := new(sync.Mutex)
	for i := 0; i < 5; i++ {
		m.Lock() // すでにロック済みならロック可能になるまで、ブロックされる
		go func(i int) {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
			m.Unlock() // ロック解除
			c <- true
		}(i)
	}
	for i := 0; i < 5; i++ {
		<-c
	}

	fmt.Println("finish!")
}

func doInCaseOnce() {
	fmt.Println("doInCaseOnce start...")

	var once sync.Once
	onceBody := func() {
		fmt.Println("Hello, World!")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody) // 1回しか呼ばれない
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	fmt.Println("finish!")
}

// waitGroupは、複数のgoroutineを待つ際に使う
func doInCaseWaitGroup() {
	fmt.Println("doInCaseWaitGroup start...")

	var wg sync.WaitGroup
	var urls = []string{"http://www.golang.org/", "http://www.google.com/", "http://www.somestupidname.com/"}

	for _, url := range urls {
		wg.Add(1) // counterをインクリメント
		go func(url string) {
			defer wg.Done() // counterをデクリメント
			time.Sleep(1 * time.Second)
			fmt.Println(url)
		}(url)
	}
	wg.Wait() // counterが0になるまで、ブロックする

	fmt.Println("finish!")
}

func doInCaseCond() {
	fmt.Println("doInCaseCond start...")

	c := make(chan bool)
	m := new(sync.Mutex)
	cond := sync.NewCond(m)

	for i := 0; i < 5; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			cond.Wait()
			fmt.Println(i)
			c <- true
		}(i)
	}
	time.Sleep(5 * time.Millisecond)
	for i := 0; i < 5; i++ {
		go func() {
			cond.Signal()
		}()
		<-c
	}

	fmt.Println("finish!")
}
