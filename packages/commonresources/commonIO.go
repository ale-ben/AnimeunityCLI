package commonresources

import (
	"io/ioutil"
	"path/filepath"
)

//WriteToFile Writes to a file
func WriteToFile(fileDir string,fileName string, fileContent string) error {
	var err error = nil

	byteArr := []byte(fileContent)
	completeFilePath := filepath.Join(fileDir,fileName)
	completeFilePath, err = filepath.Abs(completeFilePath)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(completeFilePath, byteArr, 0777)

	return err
}

//ReadFromFile Reads from a file
func ReadFromFile(fileDir string,fileName string) (string, error) {
	completeFilePath, err1 := filepath.Abs(filepath.Join(fileDir,fileName))

	if err1 != nil {
		return "", err1
	}

	data, err := ioutil.ReadFile(completeFilePath)

	return string(data),err
}