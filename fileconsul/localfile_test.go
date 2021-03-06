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

func TestToRFList(t *testing.T) {
	lfList := LFList{
		Localfile{Base: "/path/to/base", Path: "/path/to/sample1", Hash: "ac46374a846d97e22f917b6863f690ad", Data: []byte("sample1")},
		Localfile{Base: "/path/to/base", Path: "/path/to/sample2", Hash: "656b38f3402a1e8b4211fac826efd433", Data: []byte("sample2")},
	}

	ansRFList := RFList{
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample1", Hash: "ac46374a846d97e22f917b6863f690ad", Data: []byte("sample1")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample2", Hash: "656b38f3402a1e8b4211fac826efd433", Data: []byte("sample2")},
	}

	rfList := lfList.ToRFList("fileconsul")

	if !rfList.Equal(ansRFList) {
		t.Fatalf("expected result is %s, but %s", ansRFList, rfList)
	}
}

func TestSaveRemove(t *testing.T) {
	lfList := LFList{
		Localfile{Base: ".", Path: "/path/to/sample1", Hash: "ac46374a846d97e22f917b6863f690ad", Data: []byte("sample1")},
		Localfile{Base: ".", Path: "/path/to/sample2", Hash: "656b38f3402a1e8b4211fac826efd433", Data: []byte("sample2")},
	}

	err := lfList.Save()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	err = lfList.Remove()
	if err != nil {
		t.Errorf("err: %v", err)
	}
}

