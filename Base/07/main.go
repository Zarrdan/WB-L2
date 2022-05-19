package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===
Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.
Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	send := func(c <-chan interface{}) {
		for {
			select {
			case _, ok := <-c:
				if !ok {
					close(out)
					return
				}
			case _, ok := <-out:
				if !ok {
					return
				}
			}
		}
	}
	for _, c := range channels {
		go send(c)
	}
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(6*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}
