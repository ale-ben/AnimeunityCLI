package commonresources

import (
	"io/ioutil"
	"path/filepath"
)

/*
WriteToFile is used to write to a file.

Parameters:
	fileDir string //Where to create the file
	fileName string //Name of the file
	fileContent string //Content of the file
 */
func WriteToFile(fileDir string,fileName string, fileContent string) (err error) {
	byteArr := []byte(fileContent)

	completeFilePath,err := filepath.Abs(filepath.Join(fileDir,fileName))

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(completeFilePath, byteArr, 0777)

	return
}

/*
ReadFromFile is used to read from a file.

Parameters:
	fileDir string //Where to create the file
	fileName string //Name of the file

Returns	a string containing the content of the file and an error (usually == nil)
 */
func ReadFromFile(fileDir string,fileName string) (string, error) {
	completeFilePath, err := filepath.Abs(filepath.Join(fileDir,fileName))

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(completeFilePath)

	return string(data),err
}