package state

import "event-carried-state-transfer/schema"

type State struct {
	Name string
	Data schema.Person
}

type StatefulComponent interface {
	Dispatch(event schema.Event)
	AddListener(listener func(schema.Event))
	GetState() State
}

type statefulComponent struct {
	state     State
	listeners []func(schema.Event)
}

func NewStatefulComponent() StatefulComponent {
	return &statefulComponent{}
}

func (s *statefulComponent) Dispatch(event schema.Event) {
	s.state = State{Name: event.Name, Data: event.Data}
	for _, listener := range s.listeners {
		listener(event)
	}
}

func (s *statefulComponent) AddListener(listener func(schema.Event)) {
	s.listeners = append(s.listeners, listener)
}

func (s *statefulComponent) GetState() State {
	return s.state
}
