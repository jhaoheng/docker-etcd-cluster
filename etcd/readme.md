# docker image
- build image : `docker build -t orbweb/etcd:3.3.9-ubuntu16.04 .`
    - google source : `gcr.io/etcd-development/etcd`
- run : `docker run --rm -it --network etcd_network orbweb/etcd:3.3.9-ubuntu16.04 /bin/bash`
    - 確定 dnsDiscovery 已啟動
    - 啟動 : `bash /home/etcd-script/start.sh`
    - 停止 : `bash /home/etcd-script/stop.sh`
- etcd 檢查是否啟動
    - 用 etcdctl : `etcdctl member list`
    - 用 curl : `curl http://{ip}:2379/v2/members`

# etcd 範例指令
docker exec $container /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcd --version"
docker exec $container /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcdctl version"
docker exec $container /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcdctl endpoint health"
docker exec $container /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcdctl put foo bar"
docker exec $container /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcdctl get foo"

# 其他

```
--quota-backend-bytes=$((16*1024*1024)) 設定 16mb 的配額警告
--auto-compaction-retention=1 # keep one hour of history
為了解決 etcd 在壓縮歷史訊息時，產生的碎片化，並須定期執行 `etcdctl defrag` 來整理有效空間
`ETCDCTL_API=3 etcdctl endpoint status -w table` : 檢查版本數量
```