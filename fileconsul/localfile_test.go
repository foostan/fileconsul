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

func TestToMFList(t *testing.T) {
	lfList := LFList{
		Localfile{Base: "/path/to/base", Path: "/path/to/sample1", Hash: "12"},
		Localfile{Base: "/path/to/base", Path: "/path/to/sample2", Hash: "34"},
	}

	ansMFList := MFList{
		Metafile{Path: "/path/to/sample1", Url: "", Hash: "12"},
		Metafile{Path: "/path/to/sample2", Url: "", Hash: "34"},
	}

	mfList := lfList.ToMFList()

	if !mfList.Equal(ansMFList) {
		t.Fatalf("expected result is %s, but %s", ansMFList, mfList)
	}
}

func TestUrlToHash(t *testing.T) {
	_, err := UrlToHash("https://raw.githubusercontent.com/foostan/fileconsul/master/demo/config/service/apache2.json")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}
