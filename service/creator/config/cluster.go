package config

func NewCluster() *Cluster {
	return &Cluster{
		ID: "",
	}
}

type Cluster struct {
	ID string
}
