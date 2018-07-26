# api

> 在 etcd 的 container 建立後，執行

## 新增

curl -X POST {whisper_ip} -d '{"ip":"192.168.1.1", "name":"192-168-1-1"}' 

## 刪除

curl -X DELETE {whisper_ip} -d '{"ip":"192.168.1.1", "name":"192-168-1-1"}' 
