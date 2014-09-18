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
	Host string
}

func (fhA *FileHash) Equal(fhB FileHash) bool {
	if fhA.Path == fhB.Path && fhA.Hash == fhB.Hash {
		return true
	}
	return false
}

func (fhA *FileHash) In(fhsB []FileHash) bool {
	for _, fhB := range fhsB {
		if fhA.Equal(fhB) {
			return true
		}
	}
	return false
}

func (fhA *FileHash) Modified(fhB FileHash) bool {
	if fhA.Path == fhB.Path && fhA.Hash != fhB.Hash {
		return true
	}
	return false
}

func (fhA *FileHash) InMod(fhsB []FileHash) bool {
	for _, fhB := range fhsB {
		if fhA.Modified(fhB) {
			return true
		}
	}
	return false
}

func DiffFileHashs(fhsA []FileHash, fhsB []FileHash) ([]FileHash, []FileHash, []FileHash) {
	addFhs := make([]FileHash, 0)
	delFhs := make([]FileHash, 0)
	modFhs := make([]FileHash, 0)

	for _, fhA := range fhsA {
		if !fhA.In(fhsB) {
			if fhA.InMod(fhsB) {
				modFhs = append(modFhs, fhA)
			} else {
				addFhs = append(addFhs, fhA)
			}
		}
	}

	for _, fhB := range fhsB {
		if !fhB.In(fhsA) && !fhB.InMod(fhsA) {
			delFhs = append(delFhs, fhB)
		}
	}

	return addFhs, delFhs, modFhs
}

func LocalFileHashs(basepath string) ([]FileHash, error) {

	fileHashs := make([]FileHash, 0)
	searchPaths := []string{basepath}

	for len(searchPaths) > 0 {
		path := searchPaths[len(searchPaths) - 1]
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

			fileHashs = append(fileHashs, FileHash{Path: relPath, Hash: hash, Host: "localhost"})
		}

		f.Close()
	}

	return fileHashs, nil
}

func RemoteFileHashs(client *Client, prefix string) ([]FileHash, error) {
	hashprefix := filepath.Join(prefix, "hash")
	hashkvpairs, err := client.ListKV(hashprefix)
	if err != nil {
		return nil, err
	}

	fileHashs := make([]FileHash, 0)
	for _, hashkvpair := range hashkvpairs {
		relPath, err := filepath.Rel(hashprefix, hashkvpair.Key)
		if err != nil {
			return nil, fmt.Errorf("Invalid path '%s': %s", hashkvpair.Key, err)
		}

		hostkvpair, err := client.GetKV(filepath.Join(prefix, "host", relPath))
		if err != nil {
			return nil, err
		}

		fileHashs = append(fileHashs, FileHash{Path: relPath, Hash: string(hashkvpair.Value), Host: string(hostkvpair.Value)})
	}

	return fileHashs, nil
}
