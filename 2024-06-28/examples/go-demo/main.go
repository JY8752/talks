package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Item struct {
	Id string
}

func NewItem(id string) (Item, error) {
	if strings.HasPrefix(id, "d-") {
		return Item{}, fmt.Errorf("specific item not starting `d-` prefix ")
	}
	return Item{Id: id}, nil
}

// type Student struct {
// 	math    int
// 	english int
// }

// type status int

// const (
// 	_ status = iota
// 	success
// 	failed
// )

var once sync.Once

func Hello(ch chan struct{}) {
	time.Sleep(time.Second)
	ch <- struct{}{}
	close(ch)
}

func main() {
	ch := make(chan struct{})
	defer func() {
		once.Do(func() {
			close(ch)
		})
	}()

	go Hello(ch)

	<-ch
}
