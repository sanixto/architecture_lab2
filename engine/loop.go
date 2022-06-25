package engine

type Handler interface {
	Post(cmd Command)
}

type Loop struct {
	q    *cmdQueue
	stop bool

	stopSignal chan struct{}
}

func (l *Loop) Start() {
	l.q = &cmdQueue{notEmpty: make(chan struct{})}
	l.stopSignal = make(chan struct{})
	go func() {
		for !l.stop || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()

}

func (l *Loop) Post(cmd Command) {
	l.q.push(cmd)
}

func (l *Loop) AwaitFinish() {
	l.Post(stopCommand{})
	<-l.stopSignal
}
