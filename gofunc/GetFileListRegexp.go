package gofunc

import (
  "fmt"
  "errors"
  "regexp"
  "io/ioutil"
  "path/filepath"
)

////////////////////////////////////////////////////////////////////////////////
// Get selected files
////////////////////////////////////////////////////////////////////////////////
func GetFileListRegexp(dirname string, pattern string) (FilePathList []string, err error) {
  logger.Debug("start")

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // Declare variables
  // var err error <- named return value
  // var flie_path_list []string <- named return value
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
  // Compile regular expression
  //----------------------------------------------------------------------------
  logger.Debug("regular expression = %v", pattern)
  re := regexp.MustCompile(pattern)

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
    if file.IsDir() != true {
      fileno++
      logger.Debug("file%06d:%v, %v", fileno, file.Name(), re.MatchString(file.Name()))
      if re.MatchString(file.Name()) {
        FilePath     = filepath.Join(dirname, file.Name())
        FilePathList = append(FilePathList, FilePath)
      }
    }
  }

  logger.Debug("end")
  return FilePathList, nil
}
