#!/bin/bash

export ETCD_IP=$(ifconfig | awk -F'[ :]+' '/Bcast/{print $4}')
export ETCD_Name=$(echo ${ETCD_IP//./-})
export DNS_Discovery_IP=$(ping -c 1 etcd.discovery | awk -F '[()]' '/PING/{print $2}')

# 向 etcd.discovery 通知移除
echo "Notify DNS server ... to remove"
curl -X DELETE etcd.discovery:8080 -d '{"name":"'$ETCD_Name'", "ip":"'$ETCD_IP'"}'
sleep 5

# 取得 etcd id, 移除
ETCD_ID=$(etcdctl member list | grep $ETCD_Name | awk '{print $1}' | sed 's/\://g') && \
  etcdctl member remove $ETCD_ID

# 備份資料
# mkdir /home/etcd_backup
# ETCDCTL_API=3 etcdctl snapshot save /home/etcd_backup/backup-$(date '+%Y-%m-%d_%H:%M:%S').db