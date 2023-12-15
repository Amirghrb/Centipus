package state

type Job struct {
	Id         int      `json:"id"`
	Workflow   Workflow `json:"workflow"`
	Created_at string   `json:"created_at"`
	Updated_at string   `json:"updated_at"`
}
