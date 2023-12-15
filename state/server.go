package state

type Server struct {
	Name        string   `json:"name"`
	Site        string   `json:"site"`
	Description string   `json:"description"`
	Tag         []string `json:"tag"`
	Ipv4        []string `json:"ipv4"`
	Ipv6        []string `json:"ipv6"`
	Hostname    string   `json:"hostname"`
	Created_at  string   `json:"created_at"`
	Updated_at  string   `json:"updated_at"`
}
