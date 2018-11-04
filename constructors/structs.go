package constructors

type Person struct {
	ID        string   `json:"id, omitempty"`
	Firstname string   `json:"firstname, omitempty"`
	Surname   string   `json:"surname, omitempty"`
	Address   *Address `json:"address, omitempty"`
}

type Address struct {
	City  string `json:"city, omitempty"`
	State string `json:"state, omitempty"`
}
