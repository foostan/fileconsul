package fileconsul

import (
	"testing"
)

func TestConstructClient(t *testing.T) {
	_, err := NewClient(&ClientConfig{
		ConsulAddr: "localhost:8500",
		ConsulDC:   "dc1",
	})
	if err != nil {
		t.Skipf("err: %v", err)
	}
}

func TestPushKVByKeyprefix(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		ConsulAddr: "localhost:8500",
		ConsulDC:   "dc1",
	})
	if err != nil {
		t.Skipf("err: %v", err)
	}

	err = client.PushKVByKeyprefix("foo/bar/bazz", "123")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}
