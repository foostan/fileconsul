package fileconsul

import (
	"fmt"
	"path/filepath"
	"crypto/md5"
	"strings"
	"os"
)

type Remotefile struct {
	Prefix string
	Path   string
	Hash   string
	Data   []byte
}

type RFList []Remotefile

type RFDiff struct {
	Add RFList
	Del RFList
	New RFList
	Old RFList
	Eq  RFList
}

func (a *Remotefile) EqFile(b Remotefile) bool {
	if a.Path == b.Path {
		return true
	}

	return false
}

func (a *Remotefile) EqVer(b Remotefile) bool {
	if a.EqFile(b) && a.Hash == b.Hash {
		return true
	}

	return false
}

func (a *Remotefile) NeVer(b Remotefile) bool {
	if a.EqFile(b) && a.Hash != b.Hash {
		return true
	}

	return false
}

func (rfListA *RFList) Include(b Remotefile) (bool, Remotefile) {
	for _, a := range *rfListA {
		if a.EqFile(b) {
			return true, a
		}
	}
	return false, Remotefile{}
}

func (rfListA *RFList) Diff(rfListB RFList) *RFDiff {
	add := make([]Remotefile, 0)
	del := make([]Remotefile, 0)
	new := make([]Remotefile, 0)
	old := make([]Remotefile, 0)
	eq := make([]Remotefile, 0)

	for _, b := range rfListB {
		ok, a := rfListA.Include(b)
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

	for _, a := range *rfListA {
		ok, b := rfListB.Include(a)
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

	return &RFDiff{
		Add: add,
		Del: del,
		New: new,
		Old: old,
		Eq:  eq,
	}
}

func (rfListA *RFList) Equal(rfListB RFList) bool {
	rfDiff := rfListA.Diff(rfListB)
	if len(rfDiff.Add) == 0 &&
		len(rfDiff.Del) == 0 &&
		len(rfDiff.New) == 0 &&
		len(rfDiff.Old) == 0 {
		return true
	}

	return false
}

func (rfList *RFList) Empty() bool {
	return len(*rfList) <= 0
}

func (client *Client) ReadRFList(prefix string) (RFList, error) {
	kvpairs, err := client.ListKV(prefix)
	if err != nil {
		return nil, err
	}

	rfList := make([]Remotefile, 0)
	for _, kvpair := range kvpairs {
		relPath, err := filepath.Rel(prefix, kvpair.Key)
		if err != nil {
			return nil, fmt.Errorf("Invalid path '%s': %s", kvpair.Key, err)
		}

		hash := fmt.Sprintf("%x", md5.Sum(kvpair.Value))

		rfList = append(rfList, Remotefile{Prefix: prefix, Path: relPath, Hash: hash, Data: kvpair.Value})
	}

	return rfList, nil
}

func (remotefile *Remotefile) ToLocalfile(base string) Localfile {
	return Localfile{
		Base: base,
		Path: remotefile.Path,
		Hash: remotefile.Hash,
		Data: remotefile.Data,
	}
}

func (rfList *RFList) ToLFList(base string) LFList {
	lfList := make([]Localfile, 0)
	for _, remotefile := range *rfList {
		lfList = append(lfList, remotefile.ToLocalfile(base))
	}
	return lfList
}

func (rfList *RFList) Filter(pattern string) (RFList, error) {
	newRFList := make([]Remotefile, 0)
	for _, remotefile := range *rfList {
		match, err := matchPath(pattern, remotefile.Path)
		if err != nil {
			return nil, err
		}
		if match {
			newRFList = append(newRFList, remotefile)
		}
	}

	return newRFList, nil
}

func matchPath(pattern string, path string) (bool, error) {
	patternList := strings.Split(pattern, string(os.PathSeparator))
	pathList := strings.Split(path, string(os.PathSeparator))
	if len(patternList) > len(pathList) {
		return false, nil
	}

	for i := 0; i < len(patternList); i++ {
		match, err := filepath.Match(patternList[i], pathList[i])
		if err != nil {
			return false, err
		}
		if !match {
			return false, nil
		}
	}

	return true, nil
}
