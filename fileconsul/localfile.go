package fileconsul

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"net/http"
	"bytes"
)

type Localfile struct {
	Base string
	Path string
	Hash string
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

			lfList = append(lfList, Localfile{Base: basepath, Path: relPath, Hash: hash})
		}

		f.Close()
	}

	return lfList, nil
}

func (localfile *Localfile) ToMetafile() Metafile {
	return Metafile{
		Path: localfile.Path,
		Hash: localfile.Hash,
	}
}

func (lfList *LFList) ToMFList() MFList {
	mfList := make([]Metafile, 0)
	for _, localfile := range *lfList {
		mfList = append(mfList, localfile.ToMetafile())
	}
	return mfList
}

func UrlToHash(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error while downloading '%s' : %s", url, err)
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	return fmt.Sprintf("%x", md5.Sum(buf.Bytes())), nil
}
