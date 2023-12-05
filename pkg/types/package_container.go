package types

type GitPackage struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"package_type"`
	Version    int64  `json:"version_count"`
	Visibility string `json:"visibility"`
	Url        string `json:"url"`

	Owner struct {
		Id     int64  `json:"id"`
		Name   string `json:"login"`
		Type   string `json:"type"`
		NodeId string `json:"node_id"`
	} `json:"owner"`
}
