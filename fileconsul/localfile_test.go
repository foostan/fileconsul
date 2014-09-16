package fileconsul

import (
	"testing"
)

func TestReadFileHashs(t *testing.T) {
	_, err := ReadFileHashs([]string{"../test/sample"})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}
