package kubernetesdclienttest

import (
	"testing"

	"github.com/giantswarm/kubernetesdclient"
)

// Test_KubernetesdClientTest_New ensures a test client can be created without
// runtime panics and that the returned client matches the expected type.
func Test_KubernetesdClientTest_New(t *testing.T) {
	client := New()

	switch c := interface{}(client).(type) {
	case *kubernetesdclient.Client:
	default:
		t.Fatalf("expected %T got %T", &kubernetesdclient.Client{}, c)
	}
}
