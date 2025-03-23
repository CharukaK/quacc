package quacc

import (
	"os"
	"path"
)

const baseDir = ".quacc"
const noteDir = "notes"

func setupBaseDir() (fp string, err error) {
	usrHome, err := os.UserHomeDir()

	if err != nil {
		return
	}

	fp = path.Join(usrHome, baseDir, noteDir)

	if _, err := os.Stat(fp); os.IsNotExist(err) {
		err := os.MkdirAll(fp, 0755)

		if err != nil {
			return "", err
		}
	}

	return
}
