package gofunc

////////////////////////////////////////////////////////////////////////////////
// CountStrInList
// Count the of the string that matches the strings in the list
////////////////////////////////////////////////////////////////////////////////
func CountStrInList(StrList []string, str string) (cnt int) {
  logger.Debug("start")

  //----------------------------------------------------------------------------
  // Define variables
  //----------------------------------------------------------------------------
  // Declare variables
  // var cnt int <- named return value

  // Set variables
  cnt = 0

  //----------------------------------------------------------------------------
  // Confiure function for recover
  //----------------------------------------------------------------------------
  defer func() {
    logger.Debug("defer start")
    r := recover()
    if r != nil {
      logger.Error("err = %v", r)
    }
    logger.Debug("defer end")
  }()

  //----------------------------------------------------------------------------
  // Check whether str is in str_list
  //----------------------------------------------------------------------------
  for _, s := range StrList {
    if str == s {
      logger.Debug("MATCH    : [%v]", s)
      cnt++
    } else {
      logger.Debug("NOT MATCH: [%v]", s)
    }
  }

  logger.Debug("start")
  return cnt
}
