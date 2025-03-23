package fileutils

import (
	"bytes"
	"fmt"
	"os"
	"path"
)

const baseDir = ".quacc"
const noteDir = "notes"

func CreateFileIfNotExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println(filePath)
		file, err := os.Create(filePath)

		if err != nil {
			return err
		}

		file.Close()
	}

	return nil
}

func GetFileContent(filePath string) (content string, err error) {
	file, err := os.Open(filePath)

	if err != nil {
		return
	}
	defer file.Close()

	buffer := bytes.NewBuffer(make([]byte, 0))

	_, err = buffer.ReadFrom(file)

	if err != nil {
		return
	}

	content = buffer.String()

	return
}

// func GenFilePath(relPath string) string {
//     if path.Dir()
// }

func SetupBaseDir() (dir string, err error) {
	usrHome, err := os.UserHomeDir()

	if err != nil {
		return
	}

	dir = path.Join(usrHome, baseDir, noteDir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)

		if err != nil {
			return "", err
		}
	}

	return
}
