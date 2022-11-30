package main

import (
	"github.com/dharan1011/dael/pkg/el"
)

func main() {
	epoll, err := el.CreateEpoll(1024)
	if err != nil {
		panic(err)
	}
	epoll.PollEvents(-1)
}
