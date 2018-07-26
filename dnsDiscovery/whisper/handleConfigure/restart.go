package handleConfigure

import (
  "io/ioutil"
  "log"
  "os/exec"
)

/*
RestartDNSMasq ...
*/
func RestartDNSMasq() {
  // cmd := exec.Command("sh", "-c", "ps aux | grep -v grep | grep dnsmasq | awk '{print $2}'")
  // execCMD(cmd)

  stopDNSMasq()
  startDNSMasq()
}

func stopDNSMasq() {
  cmdstr := "kill -9 $(ps aux | grep -v grep | grep dnsmasq | awk '{print $2}')"
  cmd := exec.Command("sh", "-c", cmdstr)
  execCMD(cmd)
}

func startDNSMasq() {
  cmdstr := "dnsmasq --hostsdir=/home/dnsmasq/hosts --conf-dir=/home/dnsmasq/confs"
  cmd := exec.Command("sh", "-c", cmdstr)
  execCMD(cmd)
}

func execCMD(cmd *exec.Cmd) {

  stdout, err := cmd.StdoutPipe()
  defer stdout.Close()
  if err != nil {
    log.Fatal(err)
  }

  if err := cmd.Start(); err != nil {
    log.Fatal(err)
  }
  opBytes, err := ioutil.ReadAll(stdout)
  if err != nil {
    log.Fatal(err)
  }
  log.Println(string(opBytes))
}
