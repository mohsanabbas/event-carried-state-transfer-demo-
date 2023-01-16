package schema

type Person struct {
	Name       string `json:"name"`
	Age        uint8  `json:"age"`
	Address    string `json:"address"`
	Occupation string `json:"occupation"`
}

type Event struct {
	Name string
	Data Person
}
