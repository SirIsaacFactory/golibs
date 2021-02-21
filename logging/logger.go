package logging

import (
  "os"
  "errors"
)

////////////////////////////////////////////////////////////////////////////////
//  Name: logging
//
//------------------------------------------------------------------------------
//  Revision History
//  Date        Version  changed contents
//  2020/02/28  v1.00    new
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// Define type logger
////////////////////////////////////////////////////////////////////////////////
type Logger struct {
  loglevel   int       // loglevel
  streamout  *os.File  // os.Stdout, os.Stderr, or nil
  logfile    *string   // log file path
  logfileptr *os.File  // log file pointer
}

////////////////////////////////////////////////////////////////////////////////
// Logger constructor
////////////////////////////////////////////////////////////////////////////////
func NewLogger() *Logger {
  var logger *Logger = new(Logger)

  logger.loglevel   = _DEFAULT_LOGLEVEL
  logger.streamout  = os.Stderr
  logger.logfile    = nil
  logger.logfileptr = nil

  return logger
}

////////////////////////////////////////////////////////////////////////////////
// Control loglevel
////////////////////////////////////////////////////////////////////////////////
func (l *Logger) SetDebugLevel() {
  l.loglevel = DEBUG
}
func (l *Logger) SetInfoLevel() {
  l.loglevel = INFO
}
func (l *Logger) SetWarningLevel() {
  l.loglevel = WARNING
}
func (l *Logger) SetErrorLevel() {
  l.loglevel = ERROR
}
func (l *Logger) SetCriticalLevel() {
  l.loglevel = CRITICAL
}
func (l *Logger) GetLogLevel() int {
  return l.loglevel
}
func (l *Logger) SetLogLevel(loglevel int) (err error) {
  err = nil

  switch(loglevel) {
  case DEBUG:
    l.loglevel = loglevel
  case INFO:
    l.loglevel = loglevel
  case WARNING:
    l.loglevel = loglevel
  case ERROR:
    l.loglevel = loglevel
  case CRITICAL:
    l.loglevel = loglevel
  default:
    l.loglevel = _DEFAULT_LOGLEVEL
    err = errors.New("loglevel must be [DEBUG|INFO|WARNING|ERROR|CRITICAL]")
  }

  return err
}

////////////////////////////////////////////////////////////////////////////////
// Control streamout
////////////////////////////////////////////////////////////////////////////////
func (l *Logger) SetStdout() {
  l.streamout = os.Stdout
}
func (l *Logger) SetStderr() {
  l.streamout = os.Stderr
}
func (l *Logger) SetNoStreamout() {
  l.streamout = nil
}
func (l *Logger) SetStreamout(streamout *os.File) {
  l.streamout = streamout
}
func (l *Logger) GetStreamout() *os.File {
  return l.streamout
}

////////////////////////////////////////////////////////////////////////////////
// Control logfile
////////////////////////////////////////////////////////////////////////////////
func (l *Logger) SetLogFile(logfile string) {
    l.logfile = &logfile
}
func (l *Logger) GetLogFile() string {
  return *l.logfile
}
func (l *Logger) CreateLogFile() (err error) {
  l.logfileptr, err = createLogFilePointer(l.loglevel, *l.logfile)
  return err
}
func (l *Logger) OpenLogFile() (err error) {
  l.logfileptr, err = openLogFilePointer(l.loglevel, *l.logfile)
  return err
}
func (l *Logger) CloseLogFile() (err error) {
  err = closeLogFilePointer(l.loglevel, l.logfileptr)
  return err
}

////////////////////////////////////////////////////////////////////////////////
// log output
////////////////////////////////////////////////////////////////////////////////
func (l *Logger) Debug(format string, msg ...interface{}) error {
  return outputLog(l.loglevel, DEBUG, 2, l.streamout, l.logfileptr, format, msg...)
}
func (l *Logger) Info(format string, msg ...interface{}) error {
  return outputLog(l.loglevel, INFO, 2, l.streamout, l.logfileptr, format, msg...)
}
func (l *Logger) Warning(format string, msg ...interface{}) error {
  return outputLog(l.loglevel, WARNING, 2, l.streamout, l.logfileptr, format, msg...)
}
func (l *Logger) Error(format string, msg ...interface{}) error {
  return outputLog(l.loglevel, ERROR, 2, l.streamout, l.logfileptr, format, msg...)
}
func (l *Logger) Critical(format string, msg ...interface{}) error {
  return outputLog(l.loglevel, CRITICAL, 2, l.streamout, l.logfileptr, format, msg...)
}
