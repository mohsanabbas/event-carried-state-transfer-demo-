package state

import (
	"testing"

	"event-carried-state-transfer/schema"
)

func TestStatefulComponent(t *testing.T) {
	// define a test table
	tests := []struct {
		name          string
		event         schema.Event
		expectedState State
		listener      func(schema.Event)
	}{
		{
			name:          "Test Dispatch",
			event:         schema.Event{Name: "TestEvent", Data: schema.Person{Name: "John Doe", Age: 30}},
			expectedState: State{Name: "TestEvent", Data: schema.Person{Name: "John Doe", Age: 30}},
		},
		{
			name:     "Test Add Listener",
			listener: func(event schema.Event) {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewStatefulComponent()
			if test.event.Name != "" {
				s.Dispatch(test.event)
				state := s.GetState()
				if state != test.expectedState {
					t.Errorf("Expected state to be %v, but got %v", test.expectedState, state)
				}
			}
			if test.listener != nil {
				s.AddListener(test.listener)
				listeners := s.GetListeners()
				if len(listeners) != 1 {
					t.Errorf("Expected one listener, but got %d", len(listeners))
				}

			}
		})
	}
}

func TestNewStatefulComponent(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create a new stateful component",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comp := NewStatefulComponent()
			if _, ok := comp.(StatefulComponent); !ok {
				t.Errorf("NewStatefulComponent() should return a StatefulComponent but returned %T", comp)
			}
		})
	}
}

func TestDispatch(t *testing.T) {
	tests := []struct {
		name   string
		event  schema.Event
		result State
	}{
		{
			name: "Dispatch an event",
			event: schema.Event{
				Name: "John",
				Data: schema.Person{
					Name: "John Doe",
					Age:  30,
				},
			},
			result: State{
				Name: "John",
				Data: schema.Person{
					Name: "John Doe",
					Age:  30,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comp := NewStatefulComponent()
			comp.Dispatch(test.event)
			state := comp.GetState()
			if state != test.result {
				t.Errorf("Dispatch() = %v, want %v", state, test.result)
			}
		})
	}
}

func TestAddListener(t *testing.T) {
	tests := []struct {
		name           string
		listener       func(schema.Event)
		listeners      []func(schema.Event)
		event          schema.Event
		listenerCalled bool
	}{
		{
			name: "Add listener",
			listener: func(event schema.Event) {
				// do something
			},
			listeners: []func(schema.Event){
				func(event schema.Event) {
					// do something
				},
			},
			event: schema.Event{
				Name: "John",
				Data: schema.Person{
					Name: "John Doe",
					Age:  30,
				},
			},
			listenerCalled: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comp := NewStatefulComponent()
			comp.AddListener(test.listener)
			comp.Dispatch(test.event)
			listeners := comp.GetListeners()
			if len(listeners) != len(test.listeners) {
				t.Errorf("AddListener() = %v, want %v", listeners, test.listeners)
			}
			if test.listenerCalled == false {
				t.Errorf("Listener not called")
			}
		})
	}
}

func TestGetState(t *testing.T) {
	tests := []struct {
		name   string
		event  schema.Event
		result State
	}{
		{
			name: "Get current state",
			event: schema.Event{
				Name: "John",
				Data: schema.Person{
					Name: "John Doe",
					Age:  30,
				},
			},
			result: State{
				Name: "John",
				Data: schema.Person{
					Name: "John Doe",
					Age:  30,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comp := NewStatefulComponent()
			comp.Dispatch(test.event)
			state := comp.GetState()
			if state != test.result {
				t.Errorf("GetState() = %v, want %v", state, test.result)
			}
		})
	}
}

func TestGetListeners(t *testing.T) {
	tests := []struct {
		name      string
		listener  func(schema.Event)
		listeners []func(schema.Event)
	}{
		{
			name: "Get listeners",
			listener: func(event schema.Event) {

			},
			listeners: []func(schema.Event){
				func(event schema.Event) {

				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comp := NewStatefulComponent()
			comp.AddListener(test.listener)
			listeners := comp.GetListeners()
			if len(listeners) != len(test.listeners) {
				t.Errorf("GetListeners() = %v, want %v", listeners, test.listeners)
			}
		})
	}
}
