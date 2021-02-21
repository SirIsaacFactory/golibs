package gofunc

import (
  "os"
  "fmt"
  "errors"
)

////////////////////////////////////////////////////////////////////////////////
// Directory file existence
////////////////////////////////////////////////////////////////////////////////
func CheckDirExistence(fullpath string) (err error) {
  logger.Debug("start")

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // Declare variables
  // var err error <- named return value

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
  // Check Directory existence
  //----------------------------------------------------------------------------
  finfo, err := os.Stat(fullpath)
  if err != nil {
    logger.Debug("[%v] does not exist. err=%v", fullpath, err)
    logger.Debug("end")
    return errors.New(fmt.Sprintf("%v", err.Error()))
  }

  if finfo.IsDir() != true {
    logger.Debug("[%v] is a file.", fullpath)
    logger.Debug("end")
    return errors.New(fmt.Sprintf("[%v] is a file.", fullpath))
  }

  logger.Debug("end")
  return nil
}
