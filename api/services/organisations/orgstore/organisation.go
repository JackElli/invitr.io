package orgstore

type Organisation struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	People []string `json:"people"`
}
