package fileconsul

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
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
		Path:   localfile.Path,
		Hash:   localfile.Hash,
		Data:   localfile.Data,
	}
}

func (lfList *LFList) ToRFList(prefix string) RFList {
	rfList := make([]Remotefile, 0)
	for _, localfile := range *lfList {
		rfList = append(rfList, localfile.ToRemotefile(prefix))
	}
	return rfList
}

func (localfile *Localfile) Save() error {
	// temporally creating
	tmpfile, err := randstr(32)
	if err != nil {
		return fmt.Errorf("Error while generating rand string : %s", err)
	}
	err = ioutil.WriteFile(tmpfile, localfile.Data, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("Error while creating tmpfile '%s' : %s", tmpfile, err)
	}

	// atomically moving
	path := filepath.Join(localfile.Base, localfile.Path)
	err = os.MkdirAll(filepath.Dir(path), os.FileMode(0755))
	if err != nil {
		return fmt.Errorf("Error while creating '%s' : %s", path, err)
	}

	err = os.Rename(tmpfile, path)
	if err != nil {
		return fmt.Errorf("Error while moving '%s' to '%s' : %s", tmpfile, path, err)
	}

	defer os.RemoveAll(filepath.Join(localfile.Base, tmpfile))

	return nil
}

func (lfList *LFList) Save() error {
	for _, localfile := range *lfList {
		err := localfile.Save()
		if err != nil {
			return err
		}
	}

	return nil
}

func (localfile *Localfile) Remove() error {
	path := filepath.Join(localfile.Base, localfile.Path)
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("Error while removing '%s' : %s", path, err)
	}

	err = RemoveAllEmpDir(filepath.Dir(path))
	if err != nil {
		return fmt.Errorf("Error while removing '%s' : %s", path, err)
	}

	return nil
}

func (lfList *LFList) Remove() error {
	for _, localfile := range *lfList {
		err := localfile.Remove()
		if err != nil {
			return nil
		}
	}

	return nil
}

func RemoveAllEmpDir(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		err := os.Remove(path)
		if err != nil {
			return fmt.Errorf("Error while removing '%s' : %s", path, err)
		}

		return RemoveAllEmpDir(filepath.Dir(path))
	}

	return nil
}

func randstr(size int) (string, error) {
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(rb), nil
}
