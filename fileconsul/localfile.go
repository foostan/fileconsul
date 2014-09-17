package fileconsul

import (
	"fmt"
	"os"
	"path/filepath"
	"crypto/md5"
)

type FileHash struct {
	Path string
	Hash string
}

func FileHashs(paths []string) ([]FileHash, error) {
	fileHashs := make([]FileHash, 0)

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		fi, err := f.Stat()
		if err != nil {
			f.Close()
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		if !fi.IsDir() {
			data := make([]byte, fi.Size())
			_, err := f.Read(data)
			if err != nil {
				return nil, fmt.Errorf("Error reading '%s': %s", path, err)
			}
			hash := fmt.Sprintf("%x", md5.Sum(data))

			fileHashs = append(fileHashs, FileHash{Path: path, Hash: hash})
			continue
		}

		contents, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		for _, fi := range contents {
			subpath := filepath.Join(path, fi.Name())
			subpathFileHashs, err := FileHashs([]string{subpath})
			if err != nil {
				return nil, err
			}

			fileHashs = append(fileHashs, subpathFileHashs...)
		}
	}

	return fileHashs, nil
}
