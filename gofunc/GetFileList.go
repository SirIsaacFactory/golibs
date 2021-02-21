package gofunc

import (
  "fmt"
  "errors"
  "io/ioutil"
  "path/filepath"
)

////////////////////////////////////////////////////////////////////////////////
// Get file list
////////////////////////////////////////////////////////////////////////////////
func GetFileList(dirname string) (FilePathList []string, err error) {
  logger.Debug("start")

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // Declare variables
  // var FilePathList []string <- named return value
  // var err error <- named return value
  var FilePath string
  var fileno int

  // Set variables
  err = nil

  //----------------------------------------------------------------------------
  // Confiure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    logger.Debug("defer start")
    r := recover()
    if r != nil {
      logger.Error("err = %v", r)
      err = fmt.Errorf("Unexpected error occured.")
    }
    logger.Debug("defer end")
  }()


  //----------------------------------------------------------------------------
  // Reading directory
  //----------------------------------------------------------------------------
  logger.Debug("Start reading directory = [%v]", dirname)
  files, err := ioutil.ReadDir(dirname)
  if err != nil {
    logger.Debug("end")
    return nil, errors.New(fmt.Sprintf("Failed to read directory[%v] err=%v", dirname, err))
  }
  logger.Debug("Succeeded in reading directory = [%v]", dirname)

  fileno = 0
  for _, file := range files {
    fileno++
    if file.IsDir() != true {
      logger.Debug("file%06d:%v", fileno, file.Name())
      FilePath     = filepath.Join(dirname, file.Name())
      FilePathList = append(FilePathList, FilePath)
    }
  }

  logger.Debug("end")
  return FilePathList, nil
}
