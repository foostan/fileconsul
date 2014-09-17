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

func LocalFileHashs(path string) ([]FileHash, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading '%s': %s", path, err)
	}

	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("Error reading '%s': %s", path, err)
	}

	fileHashs := make([]FileHash, 0)
	if fi.IsDir() {
		contents, err := f.Readdir(-1)
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		for _, fi := range contents {
			subpath := filepath.Join(path, fi.Name())
			subpathFileHashs, err := LocalFileHashs(subpath)
			if err != nil {
				return nil, err
			}
			fileHashs = append(fileHashs, subpathFileHashs...)
		}
	} else {
		data := make([]byte, fi.Size())
		_, err := f.Read(data)
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}
		hash := fmt.Sprintf("%x", md5.Sum(data))

		fileHashs = append(fileHashs, FileHash{Path: path, Hash: hash})
	}

	f.Close()
	return fileHashs, nil
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

func RemoteFileHashs(client *Client, prefix string) ([]FileHash, error) {
	kvpairs, err := client.GetKVByKeyprefix(prefix)
	if err != nil {
		return nil, err
	}

	fileHashs := make([]FileHash, 0)
	for _, kvpair := range kvpairs {
		fileHashs = append(fileHashs, FileHash{Path: kvpair.Key, Hash: string(kvpair.Value)})
	}

	return fileHashs, nil
}
