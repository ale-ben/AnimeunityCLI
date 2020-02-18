package commonresources

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
)

var(
	//File wide logger
	commonIOLog = Log.WithField("File", "commonIO.go")
)
//TODO: Test

/*
WriteToFile is used to write to a file.

Parameters:
	fileDir string //Where to create the file
	fileName string //Name of the file
	fileContent string //Content of the file
*/
func WriteToFile(fileDir string, fileName string, fileContent string) (err error) {
	commonIOLog.WithFields(logrus.Fields{
		"Directory" : fileDir,
		"File Name" : fileName,
	}).Trace("<WriteToFile>")
	byteArr := []byte(fileContent)

	completeFilePath, err := filepath.Abs(filepath.Join(fileDir, fileName))
	commonIOLog.WithField("Complete Path",completeFilePath).Debug("Complete File Path")

	if err != nil {
		return err
	}

	commonIOLog.WithField("File Content",fileContent).Debug("File Content")
	err = ioutil.WriteFile(completeFilePath, byteArr, 0777)


	commonIOLog.WithFields(logrus.Fields{
		"Directory" : fileDir,
		"File Name" : fileName,
	}).Trace("</WriteToFile>")
	return
}

/*
ReadFromFile is used to read from a file.

Parameters:
	fileDir string //Where to create the file
	fileName string //Name of the file

Returns	a string containing the content of the file and an error (usually == nil)
*/
func ReadFromFile(fileDir string, fileName string) (string, error) {
	commonIOLog.WithFields(logrus.Fields{
		"Directory" : fileDir,
		"File Name" : fileName,
	}).Trace("<ReadFromFile>")
	completeFilePath, err := filepath.Abs(filepath.Join(fileDir, fileName))
	commonIOLog.WithField("Complete Path",completeFilePath).Debug("Complete File Path")

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(completeFilePath)
	commonIOLog.WithField("File Content",string(data)).Debug("File Content")

	commonIOLog.WithFields(logrus.Fields{
		"Directory" : fileDir,
		"File Name" : fileName,
	}).Trace("</ReadFromFile>")
	return string(data), err
}
