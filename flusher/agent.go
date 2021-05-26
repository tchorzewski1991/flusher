package flusher

//go:generate mockgen -destination=mock/exporter_mock.go -package=flusherMock github.com/tchorzewski1991/flusher/flusher Exporter

type Exporter interface {
	Export() error
}

type Agent struct {
	active   bool
	exporter Exporter
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
