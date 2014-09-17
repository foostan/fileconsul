package fileconsul

import (
	"testing"
)

func TestLocalFileHashs(t *testing.T) {
	_, err := LocalFileHashs("../test/sample")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}

func TestRemoteFileHashs(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		ConsulAddr: "localhost:8500",
		ConsulDC:   "dc1",
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	_, err = RemoteFileHashs(client, "fileconsul")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}

func TestDiffFileHashs(t *testing.T) {
	fhsA := []FileHash{
		FileHash{Path: "/path/to/sample1", Hash: "12"},
		FileHash{Path: "/path/to/sample2", Hash: "12"},
		FileHash{Path: "/path/to/sample4", Hash: "78"}}

	fhsB := []FileHash{
		FileHash{Path: "/path/to/sample1", Hash: "12"},
		FileHash{Path: "/path/to/sample2", Hash: "34"},
		FileHash{Path: "/path/to/sample3", Hash: "56"}}

	resAddFhs, resDelFhs, resModFhs := DiffFileHashs(fhsA, fhsB)

	ansAddFhs := []FileHash{
		FileHash{Path: "/path/to/sample4", Hash: "78"}}
	if len(ansAddFhs) != len(resAddFhs) {
		t.Fatalf("expected result is %s, but %s", ansAddFhs, resAddFhs)
	}
	for i := 0; i < len(ansAddFhs); i++ {
		if !ansAddFhs[i].Equal(resAddFhs[i]) {
			t.Fatalf("expected result is %s, but %s", ansAddFhs, resAddFhs)
		}
	}

	ansDelFhs := []FileHash{
		FileHash{Path: "/path/to/sample3", Hash: "56"}}
	if len(ansDelFhs) != len(resDelFhs) {
		t.Fatalf("expected result is %s, but %s", ansDelFhs, resDelFhs)
	}
	for i := 0; i < len(ansDelFhs); i++ {
		if !ansDelFhs[i].Equal(resDelFhs[i]) {
			t.Fatalf("expected result is %s, but %s", ansDelFhs, resDelFhs)
		}
	}

	ansModFhs := []FileHash{
		FileHash{Path: "/path/to/sample2", Hash: "12"}}
	if len(ansModFhs) != len(resModFhs) {
		t.Fatalf("expected result is %s, but %s", ansModFhs, resModFhs)
	}
	for i := 0; i < len(ansModFhs); i++ {
		if !ansModFhs[i].Equal(resModFhs[i]) {
			t.Fatalf("expected result is %s, but %s", ansModFhs, resModFhs)
		}
	}
}


