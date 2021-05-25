package flusher

type Agent struct {
	active bool
}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) Start() error {
	a.active = true
	return nil
}

func (a *Agent) Stop() error {
	a.active = false
	return nil
}

func (a *Agent) Active() bool {
	return a.active
}
