package gofunc

import (
  "os"
  "github.com/SirIsaacFactory/golibs/logging"
)

var logger *logging.Logger = logging.NewLogger()

func InitLogging(loglevel int, streamout *os.File) {
  logger.SetLogLevel(loglevel)
  logger.SetStreamout(streamout)
}
