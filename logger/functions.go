package logging

import (
  "errors"
  "fmt"
  "io"
  "log"
  "os"
  "path/filepath"
  "runtime"
)

////////////////////////////////////////////////////////////////////////////////
// logger
////////////////////////////////////////////////////////////////////////////////
func outputLog(loglevel int,
               msglevel int,
               depth    int,
               streamout *os.File,
               logfileptr *os.File,
               format string,
               msg ...interface{}) (err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var err errror <- named return value
  var funcname string

  // Set variables
  err = nil
  funcname = "outputLog"

  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      fmt.Fprintf(os.Stderr, "%s %s: Unexpected error occured.\n", PKGNAME, funcname)
      err = fmt.Errorf("%s %s: Unexpected error occured.", PKGNAME, funcname)
    }
  }()

  //----------------------------------------------------------------------------
  // output log when message level is grator then or equals loglevel
  //----------------------------------------------------------------------------
  if msglevel >= loglevel {

    //--------------------------------------------------------------------------
    // Define variables
    //--------------------------------------------------------------------------
    pt, fullpath, lineno, _ := runtime.Caller(depth)
    filename      := filepath.Base(fullpath)
    funcname_full := runtime.FuncForPC(pt).Name()
    funcname      := filepath.Base(funcname_full)

    // Set log flags
    log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

    // Set log output destination
    if _LOGFILEPTR != nil {
      log.SetOutput(io.MultiWriter(streamout, logfileptr))
    } else {
      log.SetOutput(streamout)
    }

    // Set prefix
    switch msglevel {
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

    log.Println(funcname, filename, lineno, fmt.Sprintf(format, msg...))
  }

  return err
}


////////////////////////////////////////////////////////////////////////////////
// Get log file pointer
////////////////////////////////////////////////////////////////////////////////
func createLogFilePointer(loglevel int,
                          logfilepath string) (logfileptr *os.File, err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var logfileptr os.File <- named return value
  // var err errror         <- named return value
  var funcname   string
  var logdirname string

  // Set variables
  logfileptr = nil
  err        = nil
  funcname   = "createLogFilePointer"

  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: Unexpected error occured.", PKGNAME, funcname)
      err = fmt.Errorf("%s %s: Unexpected error occured.", PKGNAME, funcname)
    }
  }()

  //----------------------------------------------------------------------------
  // Create log file
  //----------------------------------------------------------------------------
  // Get log directory
  logdirname, _ = filepath.Split(logfilepath)

  // Check log directory existence
  if _, err = os.Stat(logdirname); err != nil {
    outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: The log directory [%s] does not exist. err=[%v]", PKGNAME, funcname, logdirname, err.Error())
    return nil, err
  }

  // Check whether the logfilepath is a directory
  logpathinfo, err := os.Stat(logfilepath)
  if err == nil {
    if logpathinfo.IsDir() == true {
      outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: The log file path [%s] is a directory.", PKGNAME, funcname, logfilepath)
      return nil, errors.New(fmt.Sprintf("%s %s: The log file path [%s] is a directory.", PKGNAME, funcname, logfilepath))
    }
  }

  // Create logfile
  logfileptr, err = os.Create(logfilepath)
  if err != nil {
    outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: Failed to create the log file [%s]. err=[%v]", PKGNAME, funcname, logfilepath, err.Error())
    return nil, err
  }

  return logfileptr, err
}


////////////////////////////////////////////////////////////////////////////////
// Open log file
////////////////////////////////////////////////////////////////////////////////
func openLogFilePointer(loglevel int,
                        logfilepath string) (logfileptr *os.File, err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var err errror <- named return value
  var funcname   string
  var logdirname string

  // Set variables
  logfileptr = nil
  err        = nil
  funcname   = "openLogfFlePointer"

  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: Unexpected error occured.", PKGNAME, funcname)
      err = fmt.Errorf("%s %s: Unexpected error occured.", PKGNAME, funcname)
    }
  }()

  //----------------------------------------------------------------------------
  // Open log file
  //----------------------------------------------------------------------------
  // Get log directory
  logdirname, _ = filepath.Split(logfilepath)

  // Check log file existence
  if _, err := os.Stat(logfilepath); err != nil {
    outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: The log directory [%s] does not exist. err=[%v]", PKGNAME, funcname, logdirname, err.Error())
    return nil, err
  }

  // Open log file
  logfileptr, err = os.OpenFile(logfilepath, os.O_RDWR|os.O_APPEND, 0666)
  if err != nil {
    outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: Failed to open the log file [%s]. err=[%v]", PKGNAME, funcname, logfilepath, err.Error())
    return nil, err
  }

  return logfileptr, err
}

////////////////////////////////////////////////////////////////////////////////
// Close log file pointer
////////////////////////////////////////////////////////////////////////////////
func closeLogFilePointer(loglevel int, logfileptr *os.File) (err error) {

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // var err errror <- named return value
  var funcname string

  // Set variables
  err      = nil
  funcname = "closeLogFilePointer"

  //----------------------------------------------------------------------------
  // Configure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    r := recover()
    if r != nil {
      outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: Unexpected error occured.", PKGNAME, funcname)
      err = fmt.Errorf("%s %s: Unexpected error occured.", PKGNAME, funcname)
    }
  }()

  //----------------------------------------------------------------------------
  // Close log file
  //----------------------------------------------------------------------------
  if logfileptr != nil {
    err = logfileptr.Close()
    if err != nil {
      outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: Failed to close the log file. err=%v", PKGNAME, funcname, err)
      return err
    }
  } else {
    outputLog(loglevel, DEBUG, 3, os.Stderr, nil, "%s %s: The log file is not open yet.", PKGNAME, funcname)
  }

  return nil
}
