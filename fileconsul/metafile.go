package fileconsul

type Metafile struct {
	Path string
	Url  string
	Hash string
}

type MFList []Metafile

type MFDiff struct {
	Add MFList
	Del MFList
	New MFList
	Old MFList
}

func (a *Metafile) EqFile(b Metafile) bool {
	if a.Path == b.Path && a.Url == b.Url {
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

func (aList *MFList) Include(b Metafile) (bool, Metafile) {
	for _, a := range *aList {
		if a.EqFile(b) {
			return true, a
		}
	}
	return false, Metafile{}
}

func (aList *MFList) Diff(bList MFList) *MFDiff {
	add := make([]Metafile, 0)
	del := make([]Metafile, 0)
	new := make([]Metafile, 0)
	old := make([]Metafile, 0)

	for _, b := range bList {
		ok, a := aList.Include(b)
		if ok {
			switch{
			case a.EqVer(b):
				// skip
			case a.NeVer(b):
				new = append(new, a)
			}
		} else {
			del = append(del, b)
		}
	}

	for _, a := range *aList {
		ok, b := bList.Include(a)
		if ok {
			switch{
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
	}
}
