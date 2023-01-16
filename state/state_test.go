package state

import (
	"testing"

	"event-carried-state-transfer/schema"
)

func TestDispatch(t *testing.T) {
	// create a new stateful component
	s := NewStatefulComponent()

	// create a test event
	testEvent := schema.Event{Name: "TestEvent", Data: schema.Person{Name: "John Doe", Age: 30}}

	// dispatch the event
	s.Dispatch(testEvent)

	// get the current state
	state := s.GetState()

	// check if the state has the correct name and data
	if state.Name != testEvent.Name {
		t.Errorf("Expected state name to be %s, but got %s", testEvent.Name, state.Name)
	}
	if state.Data != testEvent.Data {
		t.Errorf("Expected state data to be %v, but got %v", testEvent.Data, state.Data)
	}
}

func TestAddListener(t *testing.T) {
	// create a new stateful component
	s := NewStatefulComponent()

	// create a test listener function
	testListener := func(event schema.Event) {

	}

	// add the listener to the stateful component
	s.AddListener(testListener)

	// check if the listener was added to the listeners slice
	if len(s.GetListeners()) != 1 {
		t.Errorf("Expected one listener, but got %d", len(s.GetListeners()))
	}

}

func TestGetState(t *testing.T) {
	// create a new stateful component
	s := NewStatefulComponent()

	// create a test event
	testEvent := schema.Event{Name: "TestEvent", Data: schema.Person{Name: "John Doe", Age: 30}}

	// dispatch the event
	s.Dispatch(testEvent)

	// get the current state
	state := s.GetState()

	// check if the state has the correct name and data
	if state.Name != testEvent.Name {
		t.Errorf("Expected state name to be %s, but got %s", testEvent.Name, state.Name)
	}
	if state.Data != testEvent.Data {
		t.Errorf("Expected state data to be %v, but got %v", testEvent.Data, state.Data)
	}
}

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
