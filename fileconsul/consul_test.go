package fileconsul

import (
	"testing"
)

func TestConstructClient(t *testing.T) {
	_, err := NewClient(&ClientConfig{
		ConsulAddr: "nyc1.demo.consul.io:80",
		ConsulDC:   "nyc1",
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}

