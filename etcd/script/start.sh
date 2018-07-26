#!/bin/bash

export ETCD_IP=$(ifconfig | awk -F'[ :]+' '/Bcast/{print $4}')
export ETCD_Name=$(echo ${ETCD_IP//./-})

# check work
ping -c 1 "etcd.discovery" >> /dev/null
if [[ 0 -ne $? ]]; then
  # start dns setting
  echo "DNS setting ... "
  export DNS_Discovery_IP=$(ping -c 1 dns_discovery | awk -F '[()]' '/PING/{print $2}')

  # 更新 /etc/resolv.conf
  cp /etc/resolv.conf /tmp/resolv.conf
  sed -i '1i nameserver '$DNS_Discovery_IP /tmp/resolv.conf
  # echo "nameserver "$DNS_Discovery_IP >> /tmp/resolv.conf
  cp /tmp/resolv.conf /etc/resolv.conf
fi

# 取得已註冊的 etcd api endpoint
export ETCD_API_ENDPOINT=$(dig +noall +answer SRV _etcd-server._tcp.etcd.discovery | awk '{print $8}' | head -1)

# 向 etcd.discovery 報到
# 查詢是否報到過
dig +noall +answer SRV _etcd-server._tcp.etcd.discovery | grep $ETCD_Name > /dev/null 2>&1
if [[ $? -ne 0 ]]; then
  echo "Notify DNS server ... "
  curl -X POST $DNS_Discovery_IP:8080 -d '{"name":"'$ETCD_Name'", "ip":"'$ETCD_IP'"}'
  sleep 5
fi


# 判斷 ETCD_API_ENDPOINT 是否為空, 若為空, 則為第一台
rm -rf /etcd-data
curl -X GET http://$ETCD_API_ENDPOINT:2379/v2/members > /dev/null 2>&1

# echo $?
if [[ $? -ne 0 ]]; then
  echo "==== New Cluster ===="

  /usr/local/bin/etcd \
    --data-dir /etcd-data \
    --name $ETCD_Name \
    --discovery-srv etcd.discovery \
    --initial-advertise-peer-urls http://$ETCD_IP:2380 \
    --advertise-client-urls http://0.0.0.0:2379 \
    --listen-peer-urls http://$ETCD_IP:2380 \
    --listen-client-urls http://0.0.0.0:2379 \
    --initial-cluster-token tkn \
    --initial-cluster-state new
else 
  echo "==== Exist Cluster ===="

  # remove exist etcd-data
  rm -rf /etcd-data

  # 通知 cluster 新成員
  curl http://$ETCD_API_ENDPOINT:2379/v2/members -X POST -H "Content-Type: application/json" \
  -d '{"peerURLs":["http://'$ETCD_IP':2380"]}' > /dev/null 2>&1

  sleep 5

  # Start ETCD
  /usr/local/bin/etcd \
    --data-dir /etcd-data \
    --name $ETCD_Name \
    --discovery-srv etcd.discovery \
    --initial-advertise-peer-urls http://$ETCD_IP:2380 \
    --advertise-client-urls http://0.0.0.0:2379 \
    --listen-peer-urls http://$ETCD_IP:2380 \
    --listen-client-urls http://0.0.0.0:2379 \
    --initial-cluster-token tkn \
    --initial-cluster-state existing
fi



