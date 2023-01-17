# Event-Carried State Transfer (ECST) pattern

The code provided implements the Event-Carried State Transfer (ECST) pattern. ECST is a design pattern for managing state in a distributed system.

ECST is based on the idea that state changes in the system should be triggered by events, and that these events should be broadcasted to all interested parties. This ensures that all parts of the system are aware of the current state, and can react accordingly.

The `StatefulComponent` struct in the `state` package is the main component that implements the ECST pattern. 
It has a `state` field that holds the current state of the component, and a `listeners` field that holds a slice of functions that are interested in receiving events. 
The `AddListener` function can be used to register a new listener, and the `Dispatch` function can be used to dispatch an event to the `events` channel.
The `handleEvents` function listens for events on the events channel and updates the state. It also calls all the registered listeners with the event.
The `GetState` function can be used to get the current state of the component.

The Event struct is used to represent an event in the system. It has a Name field that holds the name of the event, and a Data field that holds the data associated with the event. The Person struct is an example of the data that can be associated with an event.

The main function of the code creates a new instance of `StatefulComponent`, registers a listener and dispatches an event. It also modifies the person's age, dispatches another event and then prints the current state of the component.


This pattern can be useful in systems where state changes need to be propagated to multiple parts of the system, or where multiple parts of the system need to react to state changes. It can also be useful for implementing undo/redo functionality, or for logging state changes in a system.

### Packages

- `schema` package contains the data models, like `Person` and `Event` structs.
- `state` package contains the state management functions like `NewStatefulComponent`, `Dispatch`, `AddListener` and `GetState`.
- `conversion` package contains the json marshalling code, like `MarshalIndent` function.

### How to use

- Import the packages' `schema`, `state` and `conversion
- Create a new instance of `StatefulComponent`
- Register listeners to the component using the `AddListener` function
- Dispatch events to the channel using the `Dispatch` function
- `handleEvents` function listens for events on the events channel and updates the state. It also calls all the registered listeners with the event.
- Get the current state of the component using the `GetState` function
- Use the `MarshalIndent` function to convert the state struct to json string

### Advantages

- Enables state to be propagated to multiple parts of the system
- Allows multiple parts of the system to react to state changes
- Can be used to implement undo/redo functionality
- Can be used to log state changes in a system.