package main

import (
    "fmt"
    "sync"
    "time"
)

type Cond struct {
    L sync.Locker
    q chan struct{}
}

func NewCond(l sync.Locker) *Cond {
    return &Cond{L: l, q: make(chan struct{})}
}

func (c *Cond) Wait(v int) {
    c.L.Unlock()
	fmt.Printf("Шаг 4. Горутина %d разблокировала мьютекс в методе Wait()\n", v)
    c.q <- struct{}{}
	fmt.Printf("Шаг 5. Горутина %d Отправила в канал сигнал и взяла мьютекс\n", v)
    c.L.Lock()
	fmt.Printf("Шаг 6. Горутина %d вышла из метода Wait\n", v)
}

func (c *Cond) Signal() {
	
    // ...
}

func (c *Cond) Broadcast() {
	for {
		select {
		case  <-c.q:
			fmt.Println("select")
        default:
            return

		}
	}
    // ...
}

type State struct {
    ready bool
    cond  *Cond
}

func NewState() *State {
    return &State{cond: NewCond(&sync.Mutex{})}
}

func (s *State) WaitReady(v int) {
	fmt.Printf("Шаг 2. Горутина %d взяла мьютекс WaitReady\n", v)
    s.cond.L.Lock()
	defer s.cond.L.Unlock()
	defer fmt.Printf("Горутина %d отпустила мьютекс WaitReady\n", v)
    

    for !s.ready {
		fmt.Printf("Шаг 3. Горутина %d запустила метод Wait()\n", v)
        s.cond.Wait(v)
		fmt.Println("Вышла из Wait")
    }
}

func (s *State) SetReady() {
    s.cond.L.Lock()
    defer s.cond.L.Unlock()

    s.ready = true
	fmt.Println("Запускаем Broadcast")
    s.cond.Broadcast()
}

func main() {
    s := NewState()

    go func() {
        time.Sleep(500 * time.Millisecond)
        fmt.Println("now ready")
        s.SetReady()
    }()

    for i := 0; i < 5; i++ {
        go func(v int) {
			fmt.Printf("Шай 1. Горутина %d запустилась\n", v)
            s.WaitReady(v)
            fmt.Printf("Горутина %d ready!\n", v)
        }(i)
    }
    time.Sleep(1000 * time.Millisecond)
}