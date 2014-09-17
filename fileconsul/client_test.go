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

func TestPutKVByKeyprefix(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		ConsulAddr: "localhost:8500",
		ConsulDC:   "dc1",
	})
	if err != nil {
		t.Skipf("err: %v", err)
	}

	err = client.PutKVByKeyprefix("foo/bar/bazz", "123")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}

func TestDeleteKVByKeyprefix(t *testing.T) {
	client, err := NewClient(&ClientConfig{
	ConsulAddr: "localhost:8500",
	ConsulDC:   "dc1",
})
	if err != nil {
		t.Skipf("err: %v", err)
	}

	err = client.DeleteKVByKeyprefix("foo/bar/bazz")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}
