package main

import (
	"fmt"

	"event-carried-state-transfer/conversion"
	"event-carried-state-transfer/schema"
	"event-carried-state-transfer/state"
)

func main() {
	component := state.NewStatefulComponent()

	// Register a listener
	component.AddListener(func(event schema.Event) {
		fmt.Println("Event received:", event.Name)
	})
	person := schema.Person{
		Name:       "Aloha",
		Age:        19,
		Address:    "X street moon chowk. Pk ",
		Occupation: "Streamer",
	}
	// Dispatch an event
	component.Dispatch(schema.Event{Name: "Event 1", Data: person})

	// modify person's age
	person.Age = 20
	// dispatch modification
	component.Dispatch(schema.Event{Name: "Event 2", Data: person})

	// Get the current state
	derivedState := component.GetState()
	jsonData, err := conversion.MarshalIndent(derivedState)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Current state:", jsonData)
}
