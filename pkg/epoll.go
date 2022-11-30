package pkg

import (
	"os"
	"syscall"
)

type Epoll struct {
	efd    int
	events []syscall.EpollEvent
}

func CreateEpoll(size int) (*Epoll, error) {
	efd, err := syscall.EpollCreate(size)
	if err != nil {
		return nil, os.NewSyscallError("Error creating epoll instance", err)
	}
	return &Epoll{efd: efd, events: make([]syscall.EpollEvent, size)}, nil
}

func (this *Epoll) AddEvent(fd int, event *syscall.EpollEvent) error {
	return syscall.EpollCtl(this.efd, syscall.EPOLL_CTL_ADD, fd, event)
}

func (this *Epoll) PollEvents(timeout int) ([]syscall.EpollEvent, error) {
	ne, err := syscall.EpollWait(this.efd, this.events, timeout)
	if err != nil {
		return nil, os.NewSyscallError("Error polling event", err)
	}
	return this.events[:ne], nil
}
