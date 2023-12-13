package state

type Workflow struct {
	Id          int
	Name        string
	Description string
	Works       []Work
	Created_at  string
	Updated_at  string
}

type Work struct {
	Id      int
	Script  Script
	Server  Server
	Rel     int
	Inputs  map[string]string
	Outputs map[string]string
}
