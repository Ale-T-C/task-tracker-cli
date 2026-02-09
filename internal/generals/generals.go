package generals

import (
	"errors"
	"os"
)

func CreateFile(name string, data []byte) {
	err := os.WriteFile(name, data, 0644)
	if err != nil {
		panic(err)
	}
}

func ReadFile(name string) ([]byte, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func WriteFile(name string, data []byte) bool {
	err := os.WriteFile(name, data, 0644)
	if err != nil {
		return false
	}
	return true
}
