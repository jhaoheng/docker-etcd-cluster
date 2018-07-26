package handleConfigure

import (
  // "fmt"
  "os"
  "strings"
)

/*
AddConfiure ...
*/
func AddConfiure(NAME, IP string) {
  // 1. 讀取檔案
  // 2. replace 內容
  // 3. 新增檔案

  // NAME := "192-168-1-1"
  // IP := "192.168.1.1"

  conf := readFile("./tmpConfig/tmp.conf")
  conf = replaceARG(conf, NAME, IP)
  createNewConfigure("/home/dnsmasq/confs", NAME+".conf", conf)

  host := readFile("./tmpConfig/host")
  host = replaceARG(host, NAME, IP)
  createNewConfigure("/home/dnsmasq/hosts", NAME, host)
}

func readFile(path string) (str string) {
  file, err := os.Open(path)
  if err != nil {
    // handle the error here
    return
  }
  defer file.Close()

  // get the file size
  stat, err := file.Stat()
  if err != nil {
    return
  }
  // read the file
  bs := make([]byte, stat.Size())
  _, err = file.Read(bs)
  if err != nil {
    return
  }

  str = string(bs)
  // fmt.Println(str)
  return
}

func replaceARG(str string, NAME, IP string) (tmp string) {
  tmp = strings.Replace(str, "$NAME", NAME, -1)
  tmp = strings.Replace(tmp, "$IP", IP, -1)
  // fmt.Println(tmp)
  return
}

func createNewConfigure(DIR string, NAME string, content string) {
  path := DIR + "/" + NAME
  file, err := os.Create(path)
  if err != nil {
    // handle the error here
    return
  }
  defer file.Close()

  file.WriteString(content)
}
