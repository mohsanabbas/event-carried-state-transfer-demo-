package state

import "event-carried-state-transfer/schema"

type State struct {
	Name string
	Data schema.Person
}

type StatefulComponent struct {
	state     State
	listeners []func(schema.Event)
}

func NewStatefulComponent() *StatefulComponent {
	return &StatefulComponent{}
}

func (s *StatefulComponent) Dispatch(event schema.Event) {
	s.state = State{Name: event.Name, Data: event.Data}
	for _, listener := range s.listeners {
		listener(event)
	}
}

func (s *StatefulComponent) AddListener(listener func(schema.Event)) {
	s.listeners = append(s.listeners, listener)
}

func (s *StatefulComponent) GetState() State {
	return s.state
}
