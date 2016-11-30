package config

type Cluster struct {
	ID string `json:"id"`
}

func NewCluster() *Cluster {
	return &Cluster{
		ID: "",
	}
}
