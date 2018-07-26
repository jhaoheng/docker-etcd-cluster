package main

import (
  "encoding/json"
  "fmt"
  // "html"
  "io"
  "log"
  "net/http"
  "time"
  "whisper/handleConfigure"
)

func getBody(r io.Reader, method string) bool {
  decoder := json.NewDecoder(r)
  var obj map[string]interface{}
  err := decoder.Decode(&obj)
  if err != nil {
    fmt.Println(err)
    return false
  }

  datas := make(map[string]interface{})

  if method == "POST" {
    datas["type"] = "new"
  } else if method == "DELETE" {
    datas["type"] = "del"
  } else {
    return false
  }
  if val, ok := obj["ip"]; ok {
    datas["ip"] = val
  } else {
    return false
  }
  if val, ok := obj["name"]; ok {
    datas["name"] = val
  } else {
    return false
  }
  // put to job
  jobs <- datas
  return true
}

func route(w http.ResponseWriter, r *http.Request) {
  if ok := getBody(r.Body, r.Method); ok {
    w.WriteHeader(http.StatusOK)
  } else {
    w.WriteHeader(http.StatusBadRequest)
  }
}

func server() {
  http.HandleFunc("/", route)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

var jobs chan map[string]interface{}

func main() {
  jobs = make(chan map[string]interface{}, 20)
  go Worker()
  server()
}

/*
Worker ...
*/
func Worker() {
  for {
    select {
    case job, _ := <-jobs:
      fmt.Println("發現工作 : ", job, ", 目前等待任務 : ", len(jobs))
      updateConfigure(job)
      fmt.Println("job is done : ", job)
    default:
      time.Sleep(time.Second * 1)
      // fmt.Println("等待任務")
    }
  }
}

func updateConfigure(datas map[string]interface{}) {
  ip := datas["ip"].(string)
  name := datas["name"].(string)
  if datas["type"] == "new" {
    handleConfigure.AddConfiure(name, ip)
  } else if datas["type"] == "del" {
    handleConfigure.DelConfiure(name, ip)
  }
  handleConfigure.RestartDNSMasq()
  // time.Sleep(time.Second * 5)
}
