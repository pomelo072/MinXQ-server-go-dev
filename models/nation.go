package models

type Nation struct {
	Nation string
	Star   int
	Data   interface{}
}

type Nationls struct {
	China Nation
	Other interface{}
}
