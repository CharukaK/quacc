package fileutils

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"
)

var operatingDir = ""

const baseDir = ".quacc"
const noteDir = "notes"
const noteSuffix = ".md"

// Todo:
// - If a topic doesn't exist create the file
// - If a topic file does exist and it needs a subtopic create a directory for topic and move the file

func CreateTopicIfNotExists(topicPath string) error {
	topicDir := path.Dir(GenTopicFilePath(topicPath))

	if _, err := os.Stat(topicDir); os.IsNotExist(err) {
		if err = createDir(topicDir); err != nil {
			return err
		}

	}

	fmt.Println(fmt.Sprintf("%s%s", topicDir, noteSuffix))

	if _, err := os.Stat(fmt.Sprintf("%s%s", topicDir, noteSuffix)); err == nil {
		os.Rename(fmt.Sprintf("%s.md", topicDir), fmt.Sprintf("%s/%s%s", topicDir, path.Base(topicDir), noteSuffix))
	}
	return nil
}

func GenTopicFilePath(topicPath string) string {
	filePath := fmt.Sprintf("%s/%s", GetOperatingDir(), topicPath)
	if !strings.HasSuffix(topicPath, noteSuffix) {
		filePath += noteSuffix
	}

	return filePath
}

func GetTopicListForCompletion(p string) ([]string, error) {
	result := make([]string, 0)

	testPath := path.Join(GetOperatingDir(), p)

	files, err := os.ReadDir(testPath)
	if os.IsNotExist(err) {
		testPath = path.Dir(testPath)
		files, err = os.ReadDir(testPath)
	}

	if err != nil {
		return nil, err
	}

	relPath := strings.TrimPrefix(testPath, GetOperatingDir())
	relPath = strings.TrimPrefix(relPath, string(os.PathSeparator))
	for _, file := range files {
		if file.IsDir() {
			result = append(result, path.Join(relPath, file.Name()))
		} else {
			result = append(result, path.Join(relPath, strings.TrimSuffix(file.Name(), path.Ext(file.Name()))))
		}
	}

	return result, nil
}

func ResolveTopicFilePath(topicPath string) (string, error) {
	tp := GenTopicFilePath(topicPath)

	if _, err := os.Stat(tp); os.IsNotExist(err) {
		base := path.Base(topicPath)

		tp := GenTopicFilePath(fmt.Sprintf("%s/%s", topicPath, base))

		if _, err := os.Stat(tp); os.IsNotExist(err) {
			return "", err
		}

		return tp, nil
	}

	return tp, nil
}

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

func GetOperatingDir() string {
	return operatingDir
}

// func GenFilePath(relPath string) string {
//     if path.Dir()
// }

func SetupBaseDir() (err error) {
	usrHome, err := os.UserHomeDir()
	if err != nil {
		return
	}

	dir := path.Join(usrHome, baseDir, noteDir)

	if err = createDir(dir); err != nil {
		return err
	}

	operatingDir = dir

	return
}

func createDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)

		if err != nil {
			return err
		}
	}

	return nil
}
