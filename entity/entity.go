package entity

type Record struct {
	Origin      string
	Destination string
	Value       string
}

type CheapestRoute struct {
	Path  string `json:"path"`
	Value int    `json:"value"`
}
