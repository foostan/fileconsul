package fileconsul

import (
	"testing"
)

func TestLocalFileHashs(t *testing.T) {
	_, err := LocalFileHashs([]string{"../test/sample"})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}

func TestDiffFileHashs(t *testing.T) {
	fhsA := []FileHash{
		FileHash{Path: "/path/to/sample1", Hash: "123"},
		FileHash{Path: "/path/to/sample2", Hash: "123"}}

	fhsB := []FileHash{
		FileHash{Path: "/path/to/sample1", Hash: "123"},
		FileHash{Path: "/path/to/sample2", Hash: "456"},
		FileHash{Path: "/path/to/sample3", Hash: "789"}}

	ansDiffFhs := []FileHash{
		FileHash{Path: "/path/to/sample2", Hash: "456"},
		FileHash{Path: "/path/to/sample3", Hash: "789"}}

	resDiffFhs, err := DiffFileHashs(fhsA, fhsB)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if len(ansDiffFhs) != len(ansDiffFhs) {
		t.Fatalf("expected result is %s, but %s", ansDiffFhs, resDiffFhs)
	}

	for i := 0; i < len(ansDiffFhs); i++ {
		if !ansDiffFhs[i].Compare(resDiffFhs[i]) {
			t.Fatalf("expected result is %s, but %s", ansDiffFhs, resDiffFhs)
		}
	}
}
