package utils

import (
	"os"
	"strings"
)

func AmendFile(dir string, f func([]byte) []byte) error {
	dir = strings.ReplaceAll(dir, `\`, "/") // handle windows path
	rd, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, fi := range rd {
		if fi.IsDir() {
			continue
		}

		fullName := dir + "/" + fi.Name()
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
