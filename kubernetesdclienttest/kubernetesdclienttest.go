package kubernetesdclienttest

import "github.com/giantswarm/kubernetesdclient"

func New() *kubernetesdclient.Client {
	config := kubernetesdclient.DefaultConfig()
	client, err := kubernetesdclient.New(config)
	if err != nil {
		panic(err)
	}

	return client
}
