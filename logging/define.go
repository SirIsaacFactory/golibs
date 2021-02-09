////////////////////////////////////////////////////////////////////////////////
// This software is released under the MIT License see LICENSE.txt
// Package name: logging
// Overview    : Display log/Write log to a file.
//               When you share these properties below with other packages,
//               please use logging and its functions.
//                 loglevel   [DEBUG|INFO|WARNING|ERROR|CRITICAL]
//                 streamout  the place the logs are showed [os.Stdout|os.Stderr|nil]
//                 logfile    the log file's path
//                 logfileptr log file's file pointer
//               When you use several loggers
//               please make loggers by using this function below
//                 logger := logging.NewLogger()
//------------------------------------------------------------------------------
// Author: Isaac Factory (sir.isaac.factory@icloud.com)
// Repository: https://github.com/SirIsaacFactory/golibs
// Date: 2021/02/09
// Code version: v1.00
////////////////////////////////////////////////////////////////////////////////
package logging

const (
  VERSION = "v1.00"
  PKGNAME = "logging"
)

////////////////////////////////////////////////////////////////////////////////
// Define constant variables
////////////////////////////////////////////////////////////////////////////////
// loglevels
const (
  DEBUG = iota
  INFO
  WARNING
  ERROR
  CRITICAL
)

// Default loglevel
const _DEFAULT_LOGLEVEL = ERROR
