package fileconsul

import (
	"testing"
)

func TestFileHashs(t *testing.T) {
	_, err := FileHashs([]string{"../test/sample"})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}
