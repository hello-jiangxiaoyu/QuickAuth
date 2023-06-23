package utils

import (
	"os"
	"strings"
)

func AmendFile(path string, f func([]byte) []byte) error {
	path = strings.ReplaceAll(path, `\`, "/") // handle windows path
	rd, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, fi := range rd {
		if fi.IsDir() {
			continue
		}

		fullName := path + "/" + fi.Name()
		src, err := os.ReadFile(fullName)
		if err != nil {
			return err
		}
		amend := f(src)
		if err = os.WriteFile(fullName, amend, 666); err != nil {
			return err
		}
	}
	return nil
}
