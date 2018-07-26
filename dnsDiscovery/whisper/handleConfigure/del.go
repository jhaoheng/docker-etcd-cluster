package handleConfigure

import (
  "fmt"
  "os"
)

/*
DelConfiure ...
*/
func DelConfiure(NAME, IP string) {
  // 1. 刪除掉 conf/host
  delFile("/home/dnsmasq/confs", NAME+".conf")
  delFile("/home/dnsmasq/hosts", NAME)
}

func delFile(DIR, NAME string) {
  path := DIR + "/" + NAME
  // delete file
  var err = os.Remove(path)
  if isError(err) {
    return
  }

  fmt.Println("==> done deleting file : ", path)
}

func isError(err error) bool {
  if err != nil {
    fmt.Println(err.Error())
  }

  return (err != nil)
}
