package fileconsul

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
)

type Localfile struct {
	Base string
	Path string
	Hash string
	Data []byte
}

type LFList []Localfile

func ReadLFList(basepath string) (LFList, error) {
	lfList := make([]Localfile, 0)
	searchPaths := []string{basepath}

	for len(searchPaths) > 0 {
		path := searchPaths[len(searchPaths)-1]
		searchPaths = searchPaths[:len(searchPaths)-1]

		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		fi, err := f.Stat()
		if err != nil {
			f.Close()
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		if fi.IsDir() {
			contents, err := f.Readdir(-1)
			if err != nil {
				return nil, fmt.Errorf("Error reading '%s': %s", path, err)
			}

			for _, fi := range contents {
				subpath := filepath.Join(path, fi.Name())
				searchPaths = append(searchPaths, subpath)
			}
		} else {
			data := make([]byte, fi.Size())
			_, err := f.Read(data)
			if err != nil {
				return nil, fmt.Errorf("Error reading '%s': %s", path, err)
			}

			relPath, err := filepath.Rel(basepath, path)
			if err != nil {
				return nil, fmt.Errorf("Invalid path '%s': %s", path, err)
			}

			hash := fmt.Sprintf("%x", md5.Sum(data))

			lfList = append(lfList, Localfile{Base: basepath, Path: relPath, Hash: hash, Data: data})
		}

		f.Close()
	}

	return lfList, nil
}

func (localfile *Localfile) ToRemotefile(prefix string) Remotefile {
	return Remotefile{
		Prefix: prefix,
		Path: localfile.Path,
		Hash: localfile.Hash,
		Data: localfile.Data,
	}
}

func (lfList *LFList) ToRFList(prefix string) RFList {
	rfList := make([]Remotefile, 0)
	for _, localfile := range *lfList {
		rfList = append(rfList, localfile.ToRemotefile(prefix))
	}
	return rfList
}
