package logger

import (
  "os"
  "io"
  "fmt"
  "log"
  "runtime"
  "errors"
  "path/filepath"
)

////////////////////////////////////////////////////////////////////////////////
//  Name: logger
//
//------------------------------------------------------------------------------
//  Revision History
//  Date        Version  changed contents
//  2020/02/28  v1.00    new
////////////////////////////////////////////////////////////////////////////////
const (
  VERSION = "v1.00"
)

////////////////////////////////////////////////////////////////////////////////
// Loglevel
////////////////////////////////////////////////////////////////////////////////
const (
  DEBUG = iota
  INFO
  WARNING
  ERROR
  CRITICAL
)


////////////////////////////////////////////////////////////////////////////////
// Define variables
////////////////////////////////////////////////////////////////////////////////
// Loglevel(Default:DEBUG)
var _LOGLEVEL = INFO
// Logfile
var _LOGFILE *os.File


////////////////////////////////////////////////////////////////////////////////
// logger
////////////////////////////////////////////////////////////////////////////////
func Msg(loglevel int, format string, values ...interface{}) (err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var err errror <- named return value
  err = nil

  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      fmt.Println("logwrt: Unexpected error occured.")
      err = fmt.Errorf("logwrt: Unexpected error occured.")
    }
  }()

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  pt, fullpath, lineno, _ := runtime.Caller(1)
  filename      := filepath.Base(fullpath)
  funcname_full := runtime.FuncForPC(pt).Name()
  funcname      := filepath.Base(funcname_full)

  // Set log output flags
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

  // set log output
  if _LOGFILE != nil {
    log.SetOutput(io.MultiWriter(_LOGFILE, os.Stdout))
  } else {
    log.SetOutput(os.Stdout)
  }

  // Set prefix
  switch loglevel {
  case DEBUG:
    log.SetPrefix("[debug]    ")
  case INFO:
    log.SetPrefix("[info]     ")
  case WARNING:
    log.SetPrefix("[warning]  ")
  case ERROR:
    log.SetPrefix("[error]    ")
  case CRITICAL:
    log.SetPrefix("[critical] ")
  default:
    log.SetPrefix("[unexpected prefix] ")
  }

  // Log write
  if loglevel >= _LOGLEVEL {
    log.Println(funcname, filename, lineno, fmt.Sprintf(format, values...))
  }

  // return
  return err
}

////////////////////////////////////////////////////////////////////////////////
// Set loglevel
////////////////////////////////////////////////////////////////////////////////
func Setloglevel(loglevel int) (err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var err errror <- named return value
  err = nil

  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      fmt.Println("logwrt: Unexpected error occured.")
      err = fmt.Errorf("logwrt: Unexpected error occured.")
    }
  }()

  //----------------------------------------------------------------------------
  // Set loglevel
  //----------------------------------------------------------------------------
  // if unkown loglevel is given, set loglevel DEBUG
  if loglevel < DEBUG || CRITICAL < loglevel {
    _LOGLEVEL = DEBUG
    Msg(ERROR, "Unexpected loglevel was given. loglevel was set to DEBUG.")
    return errors.New("Unexpected loglevel was given. loglevel was set to DEBUG.")
  } else {
    _LOGLEVEL = loglevel
  }

  // Return
  return nil
}


////////////////////////////////////////////////////////////////////////////////
// Create logfile
////////////////////////////////////////////////////////////////////////////////
func Createlogfile(logfilepath string) (err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var err errror <- named return value
  err = nil


  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      Msg(CRITICAL, "Unexpected error occured.")
      err = fmt.Errorf("Unexpected error occured.")
    }
  }()


  //----------------------------------------------------------------------------
  // Create logfile
  //----------------------------------------------------------------------------
  Msg(DEBUG, "start")

  // Get log directory
  logdirname, logfilename := filepath.Split(logfilepath)
  Msg(DEBUG, "logdirname=%s, logfilename=%s", logdirname, logfilename)

  // Check log directory existence
  if _, err := os.Stat(logdirname); err != nil {
    Msg(ERROR, "err=%v", err.Error())
    return errors.New(fmt.Sprintf("directory[%s] does not exist.", logdirname))
  }

  // Create logfile
  local_logfile, err := os.Create(logfilepath)
  if err != nil {
    Msg(ERROR, "err=%v", err.Error())
    return errors.New(fmt.Sprintf("Failed to open logfile[%s].", logfilepath))
  }
  _LOGFILE = local_logfile

  // Return
  Msg(DEBUG, "end")
  return err
}


////////////////////////////////////////////////////////////////////////////////
// Open logfile
////////////////////////////////////////////////////////////////////////////////
func Openlogfile(logfilepath string) (err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var err errror <- named return value
  err = nil


  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      Msg(CRITICAL, "Unexpected error occured.")
      err = fmt.Errorf("Unexpected error occured.")
    }
  }()


  //----------------------------------------------------------------------------
  // Open logfile
  //----------------------------------------------------------------------------
  Msg(DEBUG, "start")

  // Get log directory
  logdirname, logfilename := filepath.Split(logfilepath)
  Msg(DEBUG, "logdirname=%s, logfilename=%s", logdirname, logfilename)

  // Check log file existence
  if _, err := os.Stat(logfilepath); err != nil {
    Msg(ERROR, "err=%v", err.Error())
    return errors.New(fmt.Sprintf("logfile[%s] does not exist.", logdirname))
  }

  // Open logfile
  local_logfile, err := os.OpenFile(logfilepath, os.O_RDWR|os.O_APPEND, 0666)
  if err != nil {
    Msg(ERROR, "err=%v", err.Error())
    return errors.New(fmt.Sprintf("Failed to open logfile[%s].", logfilepath))
  }

  _LOGFILE = local_logfile

  // Return
  return err
}


////////////////////////////////////////////////////////////////////////////////
// Close logfile
////////////////////////////////////////////////////////////////////////////////
func Closelogfile() (err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var err errror <- named return value
  err = nil


  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      Msg(CRITICAL, "Unexpected error occured.")
      err = fmt.Errorf("Unexpected error occured.")
    }
  }()


  //----------------------------------------------------------------------------
  // Close logfile
  //----------------------------------------------------------------------------
  Msg(DEBUG, "start")

  if _LOGFILE != nil {
    Msg(DEBUG, "close logfile.")
    _LOGFILE.Close()
  } else {
    Msg(DEBUG, "logfile is not opened yet.")
  }

  // Return
  Msg(DEBUG, "end")
  return err
}
