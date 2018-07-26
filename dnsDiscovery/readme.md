> http://www.thekelleys.org.uk/dnsmasq/docs/dnsmasq-man.html

# How to run dnsDiscovery
- image : `docker build -t orbweb/etcd_dns_discovery:1.0 .`
- Run : `docker run --rm --cap-add NET_ADMIN -p 8080:8080 --hostname etcd.discovery orbweb/etcd_dns_discovery:1.0`

# dnsDiscovery 運作邏輯
- 當 whisper api 被呼叫時 (post / delete), dns server 會變更 dnsmasq 的 configure, 並且重新啟動 dnsmasq

# Structure
- dnsmasq, 負責 dns server
    - 執行指令 : `dnsmasq --hostsdir=/home/dnsmasq/hosts --conf-dir=/home/dnsmasq/confs`
    - 啟動時注意 : 
        - hostname : etcd.discovery
        - capability (--cap-add) : NET_ADMIN
    - 設定檔 : conf / host
        - host 的格式 : `<container_ip> <name>.etcd.discovery`
        - conf 的格式 : `<_service}.<_prot>.[<domain>],[<target>[,<port>[,<priority>[,<weight>]]]]`
    - 如何手動更新 dnsmasq 的設定
        1. `kill -9 $(ps aux | grep -v grep | grep dnsmasq | awk '{print $2}')`
        2. `dnsmasq --hostsdir=/home/dnsmasq/hosts --conf-dir=/home/dnsmasq/confs`
    - 如何驗證設定成功
        - `dig +noall +answer SRV _etcd-server._tcp.etcd.discovery`
        - `dig +noall +answer {name}.etcd.discovery`
- whisper, 負責更新 dns server 的 config, 功能有
    1. 接收 新增/刪除 的資訊, 更新 config, 重新啟動 dnsmasq
    2. 透過 goroutine 的功能, 鎖住進程, 避免同時更新 config, 造成儲存時的錯誤
    3. 重新啟動 dnsmasq
    4. api 請參考 `./whisper/readme.md`



