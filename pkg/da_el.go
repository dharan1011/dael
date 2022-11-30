package pkg

type DaEL struct {
	epoll      *Epoll
	fileEvents map[int]*DaFileEvent
	Name       string
}

func NewDaEventLoop(size int) (*DaEL, error) {
	epoll, err := CreateEpoll(size)
	if err != nil {
		return nil, err
	}
	return &DaEL{
		epoll:      epoll,
		fileEvents: make(map[int]*DaFileEvent),
		Name:       "epoll",
	}, nil
}

func (this *DaEL) MointorFileEvent(event *DaFileEvent) error {
	this.fileEvents[event.fd] = event
	return this.epoll.AddEvent(event.fd, event.ToNative())
}

func (this *DaEL) PollEvents(timeout int) ([]DaFileEvent, error) {
	var op []DaFileEvent
	if events, err := this.epoll.PollEvents(0); err != nil {
		return nil, err
	} else {
		for _, e := range events {
			op = append(op, *this.fileEvents[int(e.Fd)])
		}
	}
	return op, nil
}
