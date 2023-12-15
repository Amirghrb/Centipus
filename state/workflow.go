package state

type Workflow struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Works       []Work `json:"works"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type Work struct {
	Id      int               `json:"id"`
	Script  Script            `json:"script"`
	Server  Server            `json:"server"`
	Rel     int               `json:"rel"`
	Inputs  map[string]string `json:"inputs"`
	Outputs map[string]string `json:"outputs"`
}
