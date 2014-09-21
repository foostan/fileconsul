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
		Localfile{Base: "/path/to/base", Path: "/path/to/sample1", Hash: "12"},
		Localfile{Base: "/path/to/base", Path: "/path/to/sample2", Hash: "34"},
	}

	ansRFList := RFList{
		Remotefile{Path: "/path/to/sample1", Url: "", Hash: "12"},
		Remotefile{Path: "/path/to/sample2", Url: "", Hash: "34"},
	}

	rfList := lfList.ToRFList()

	if !rfList.Equal(ansRFList) {
		t.Fatalf("expected result is %s, but %s", ansRFList, rfList)
	}
}

func TestUrlToHash(t *testing.T) {
	_, err := UrlToHash("https://raw.githubusercontent.com/foostan/fileconsul/master/demo/config/service/apache2.json")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}
