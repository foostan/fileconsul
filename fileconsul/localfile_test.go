package fileconsul

import (
	"testing"
)

func TestReadLFList(t *testing.T) {
	_, err := ReadLFList("..")
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
