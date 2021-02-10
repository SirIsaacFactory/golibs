package logging

import (
  "os"
)

////////////////////////////////////////////////////////////////////////////////
// Define package variables
////////////////////////////////////////////////////////////////////////////////
var _LOGLEVEL   int      = _DEFAULT_LOGLEVEL   // loglevel
var _STREAMOUT  *os.File = os.Stderr           // os.Stdout, os.Stderr, or nil
var _LOGFILE    *string  = nil                 // log file path
var _LOGFILEPTR *os.File = nil                 // log file pointer

////////////////////////////////////////////////////////////////////////////////
// Set loglevel
////////////////////////////////////////////////////////////////////////////////
func SetDebugLevel() {
  _LOGLEVEL = DEBUG
}
func SetInfoLevel() {
  _LOGLEVEL = INFO
}
func SetWarningLevel() {
  _LOGLEVEL = WARNING
}
func SetErrorLevel() {
  _LOGLEVEL = ERROR
}
func SetCriticalLevel() {
  _LOGLEVEL = CRITICAL
}
func GetLogLevel() int {
  return _LOGLEVEL
}

////////////////////////////////////////////////////////////////////////////////
// Set streamout
///////////////////////////////////////////////////////////////////////////////
func SetStdout() {
  _STREAMOUT = os.Stdout
}
func SetStderr() {
  _STREAMOUT = os.Stderr
}
func SetNoStreamout() {
  _STREAMOUT = nil
}
func SetStreamout(streamout *os.File) {
  _STREAMOUT = streamout
}
func GetStreamout() *os.File {
  return _STREAMOUT
}

////////////////////////////////////////////////////////////////////////////////
// Set log file
////////////////////////////////////////////////////////////////////////////////
func SetLogFile(logfile string) {
  _LOGFILE = &logfile
}
func GetLogFile() string {
  return *_LOGFILE
}
func CreateLogFile() (err error) {
  _LOGFILEPTR, err = createLogFilePointer(_LOGLEVEL, *_LOGFILE)
  return err
}
func OpenLogFile() (err error) {
  _LOGFILEPTR, err = openLogFilePointer(_LOGLEVEL, *_LOGFILE)
  return err
}
func CloseLogFile() (err error) {
  err = closeLogFilePointer(_LOGLEVEL, _LOGFILEPTR)
  return err
}

////////////////////////////////////////////////////////////////////////////////
// log output
////////////////////////////////////////////////////////////////////////////////
func Debug(format string, msg ...interface{}) error {
  return outputLog(_LOGLEVEL, DEBUG, 2, _STREAMOUT, _LOGFILEPTR, format, msg...)
}
func Info(format string, msg ...interface{}) error {
  return outputLog(_LOGLEVEL, INFO, 2, _STREAMOUT, _LOGFILEPTR, format, msg...)
}
func Warning(format string, msg ...interface{}) error {
  return outputLog(_LOGLEVEL, WARNING, 2, _STREAMOUT, _LOGFILEPTR, format, msg...)
}
func Error(format string, msg ...interface{}) error {
  return outputLog(_LOGLEVEL, ERROR, 2, _STREAMOUT, _LOGFILEPTR, format, msg...)
}
func Critical(format string, msg ...interface{}) error {
  return outputLog(_LOGLEVEL, CRITICAL, 2, _STREAMOUT, _LOGFILEPTR, format, msg...)
}
