package fileconsul

import (
	"strings"
	"testing"
)

func TestDiff(t *testing.T) {
	NewMFList := MFList{
		Metafile{Path: "/path/to/sample1", Url: "http://path/to/sample1", Hash: "12"},
		Metafile{Path: "/path/to/sample2", Url: "http://path/to/sample2", Hash: "12"},
		Metafile{Path: "/path/to/sample5", Url: "http://path/to/sample4", Hash: "90"},
		Metafile{Path: "/path/to/sample6", Url: "http://path/to/sample4", Hash: "12"},
	}

	OldMFList := MFList{
		Metafile{Path: "/path/to/sample1", Url: "http://path/to/sample1", Hash: "12"},
		Metafile{Path: "/path/to/sample2", Url: "http://path/to/sample2", Hash: "34"},
		Metafile{Path: "/path/to/sample3", Url: "http://path/to/sample3", Hash: "56"},
		Metafile{Path: "/path/to/sample4", Url: "http://path/to/sample4", Hash: "78"},
	}

	mfDiff := NewMFList.Diff(OldMFList)

	addMfDiff := MFList{
		Metafile{Path: "/path/to/sample5", Url: "http://path/to/sample4", Hash: "90"},
		Metafile{Path: "/path/to/sample6", Url: "http://path/to/sample4", Hash: "12"}}
	if len(mfDiff.Add) != len(addMfDiff) {
		t.Fatalf("expected result is %s, but %s", addMfDiff, mfDiff.Add)
	}
	for i := 0; i < len(addMfDiff); i++ {
		if !addMfDiff[i].EqVer(mfDiff.Add[i]) {
			t.Fatalf("expected result is %s, but %s", addMfDiff[i], mfDiff.Add[i])
		}
	}

	delMfDiff := MFList{
		Metafile{Path: "/path/to/sample3", Url: "http://path/to/sample3", Hash: "56"},
		Metafile{Path: "/path/to/sample4", Url: "http://path/to/sample4", Hash: "78"}}
	if len(mfDiff.Del) != len(delMfDiff) {
		t.Fatalf("expected result is %s, but %s", delMfDiff, mfDiff.Del)
	}
	for i := 0; i < len(delMfDiff); i++ {
		if !delMfDiff[i].EqVer(mfDiff.Del[i]) {
			t.Fatalf("expected result is %s, but %s", delMfDiff[i], mfDiff.Del[i])
		}
	}

	newMfDiff := MFList{
		Metafile{Path: "/path/to/sample2", Url: "http://path/to/sample2", Hash: "12"}}
	if len(mfDiff.New) != len(newMfDiff) {
		t.Fatalf("expected result is %s, but %s", newMfDiff, mfDiff.New)
	}
	for i := 0; i < len(newMfDiff); i++ {
		if !newMfDiff[i].EqVer(mfDiff.New[i]) {
			t.Fatalf("expected result is %s, but %s", newMfDiff[i], mfDiff.New[i])
		}
	}

	oldMfDiff := MFList{
		Metafile{Path: "/path/to/sample2", Url: "http://path/to/sample2", Hash: "34"}}
	if len(mfDiff.Old) != len(oldMfDiff) {
		t.Fatalf("expected result is %s, but %s", oldMfDiff, mfDiff.Old)
	}
	for i := 0; i < len(oldMfDiff); i++ {
		if !oldMfDiff[i].EqVer(mfDiff.Old[i]) {
			t.Fatalf("expected result is %s, but %s", oldMfDiff[i], mfDiff.Old[i])
		}
	}
}

func TestReadMFList(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		ConsulAddr: "localhost:8500",
		ConsulDC:   "dc1",
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	_, err = client.ReadMFList("fileconsul")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}

func TestStrToMFValue(t *testing.T) {
	url := "http://path/to/sample1"
	hash := "12"
	value := strings.Join([]string{url, hash}, ",")

	mfValue := StrToMFValue(value)

	if mfValue.Url != url {
		t.Fatalf("expected result is %s, but %s", url, mfValue.Url)
	}

	if mfValue.Hash != hash {
		t.Fatalf("expected result is %s, but %s", hash, mfValue.Hash)
	}
}
