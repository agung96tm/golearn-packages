package lib

import (
	"fmt"
	"io"
	"os"
)

type Storage struct {
}

func NewStorage() Storage {
	return Storage{}
}

func (s *Storage) Put(fileName string, reader io.Reader) (string, error) {
	checkfile := fmt.Sprintf("%s/%s", "media", fileName)

	out, err := os.Create("ui/" + checkfile)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, reader)
	if err != nil {
	}

	return checkfile, nil
}

func (s *Storage) Get(fileName string) (string, error) {
	return "", nil
}

func (s *Storage) Delete(fullPath string) error {
	err := os.Remove("ui/" + fullPath)
	if err != nil {
		return err
	}
	return nil
}
