package fileconsul

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
)

type FileHash struct {
	Path string
	Hash string
}

func (fhA *FileHash) Compare(fhB FileHash) bool {
	if fhA.Path == fhB.Path && fhA.Hash == fhB.Hash {
		return true
	}
	return false
}

func (fhA *FileHash) In(fhsB []FileHash) bool {
	for _, fhB := range fhsB {
		if fhA.Compare(fhB) {
			return true
		}
	}
	return false
}

func DiffFileHashs(fhsA []FileHash, fhsB []FileHash) ([]FileHash, error) {
	diffFhs := make([]FileHash, 0)

	for _, fhB := range fhsB {
		if !fhB.In(fhsA) {
			diffFhs = append(diffFhs, fhB)
		}
	}

	return diffFhs, nil
}

func LocalFileHashs(basepath string) ([]FileHash, error) {

	fileHashs := make([]FileHash, 0)
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

			fileHashs = append(fileHashs, FileHash{Path: relPath, Hash: hash})
		}

		f.Close()
	}

	return fileHashs, nil
}

func RemoteFileHashs(client *Client, prefix string) ([]FileHash, error) {
	kvpairs, err := client.GetKVByKeyprefix(prefix)
	if err != nil {
		return nil, err
	}

	fileHashs := make([]FileHash, 0)
	for _, kvpair := range kvpairs {
		relPath, err := filepath.Rel(prefix, kvpair.Key)
		if err != nil {
			return nil, fmt.Errorf("Invalid path '%s': %s", kvpair.Key, err)
		}

		fileHashs = append(fileHashs, FileHash{Path: relPath, Hash: string(kvpair.Value)})
	}

	return fileHashs, nil
}
