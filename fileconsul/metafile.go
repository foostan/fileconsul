package fileconsul

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Metafile struct {
	Prefix string
	Path   string
	Url    string
	Hash   string
}

type MFList []Metafile

type MFDiff struct {
	Add MFList
	Del MFList
	New MFList
	Old MFList
	Eq  MFList
}

type MFValue struct {
	Url  string
	Hash string
}

func (a *Metafile) EqFile(b Metafile) bool {
	if a.Path == b.Path &&
		(a.Url == "" || b.Url == "" || a.Url == b.Url) {
		return true
	}

	return false
}

func (a *Metafile) EqVer(b Metafile) bool {
	if a.EqFile(b) && a.Hash == b.Hash {
		return true
	}

	return false
}

func (a *Metafile) NeVer(b Metafile) bool {
	if a.EqFile(b) && a.Hash != b.Hash {
		return true
	}

	return false
}

func (mfListA *MFList) Include(b Metafile) (bool, Metafile) {
	for _, a := range *mfListA {
		if a.EqFile(b) {
			return true, a
		}
	}
	return false, Metafile{}
}

func (mfListA *MFList) Diff(mfListB MFList) *MFDiff {
	add := make([]Metafile, 0)
	del := make([]Metafile, 0)
	new := make([]Metafile, 0)
	old := make([]Metafile, 0)
	eq := make([]Metafile, 0)

	for _, b := range mfListB {
		ok, a := mfListA.Include(b)
		if ok {
			switch {
			case a.EqVer(b):
				eq = append(eq, a)
			case a.NeVer(b):
				new = append(new, a)
			}
		} else {
			del = append(del, b)
		}
	}

	for _, a := range *mfListA {
		ok, b := mfListB.Include(a)
		if ok {
			switch {
			case b.EqVer(a):
				// skip
			case b.NeVer(a):
				old = append(old, b)
			}
		} else {
			add = append(add, a)
		}
	}

	return &MFDiff{
		Add: add,
		Del: del,
		New: new,
		Old: old,
		Eq:  eq,
	}
}

func (mfListA *MFList) Equal(mfListB MFList) bool {
	mfDiff := mfListA.Diff(mfListB)
	if len(mfDiff.Add) == 0 &&
		len(mfDiff.Del) == 0 &&
		len(mfDiff.New) == 0 &&
		len(mfDiff.Old) == 0 {
		return true
	}

	return false
}

func (client *Client) ReadMFList(prefix string) (MFList, error) {
	kvpairs, err := client.ListKV(prefix)
	if err != nil {
		return nil, err
	}

	mfList := make([]Metafile, 0)
	for _, kvpair := range kvpairs {
		relPath, err := filepath.Rel(prefix, kvpair.Key)
		if err != nil {
			return nil, fmt.Errorf("Invalid path '%s': %s", kvpair.Key, err)
		}

		mfValue := StrToMFValue(string(kvpair.Value))

		mfList = append(mfList, Metafile{Prefix: prefix, Path: relPath, Url: mfValue.Url, Hash: mfValue.Hash})
	}

	return mfList, nil
}

func StrToMFValue(str string) MFValue {
	splited := strings.Split(str, ",")
	return MFValue{
		Url:  splited[0],
		Hash: splited[1],
	}
}

func (mfValue *MFValue) ToStr() string {
	return strings.Join([]string{mfValue.Url, mfValue.Hash}, ",")
}
