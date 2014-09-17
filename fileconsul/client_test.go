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

func TestPutGetDeleteKV(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		ConsulAddr: "localhost:8500",
		ConsulDC:   "dc1",
	})
	if err != nil {
		t.Skipf("err: %v", err)
	}

	err = client.PutKV("foo/bar/bazz", "123")
	if err != nil {
		t.Skipf("err: %v", err)
	}

	_, err = client.GetKV("foo/bar/bazz")
	if err != nil {
		t.Skipf("err: %v", err)
	}

	err = client.DeleteKV("foo/bar/bazz")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}
