package el

import (
	"os"
	"syscall"
)

type EventHandler func(interface{})

type daEvent int

const (
	DA_READABLE_EVENT = (1 << 1)
	DA_WRITABLE_EVENT = (1 << 2)
)

type DaFileEvent struct {
	fd      int
	event   daEvent
	handler EventHandler
}

func (this *DaFileEvent) ToNative() *syscall.EpollEvent {
	var mask uint32 = 0
	if this.event&DA_READABLE_EVENT != 0 {
		mask |= syscall.EPOLLIN
	}
	if this.event&DA_WRITABLE_EVENT != 0 {
		mask |= syscall.EPOLLOUT
	}
	return &syscall.EpollEvent{Fd: int32(this.fd), Events: mask}
}

func NewDaFileEvent(fd int, option daEvent, handler EventHandler) (*DaFileEvent, error) {
	if err := SetNonBlocking(fd); err != nil {
		return nil, err
	}
	return &DaFileEvent{fd: fd, event: option, handler: handler}, nil
}

func SetNonBlocking(fd int) error {
	err := syscall.SetNonblock(fd, true)
	if err != nil {
		return os.NewSyscallError("Setting non blocking operation", err)
	}
	return nil
}
