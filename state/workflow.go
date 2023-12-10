package state

type workflow struct {
	id          int
	name        string
	description string
	works       []work
	created_at  string
	updated_at  string
}

type work struct {
	id      int
	script  script
	server  server
	rel     int
	inputs  map[string]string
	outputs map[string]string
}
