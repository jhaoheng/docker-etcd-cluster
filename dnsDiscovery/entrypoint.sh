#!/bin/bash

# run dnsmasq
cp /etc/resolv.conf /tmp/resolv.conf
sed -i '1i nameserver 127.0.0.1' /tmp/resolv.conf
cp /tmp/resolv.conf /etc/resolv.conf
dnsmasq --hostsdir=/home/dnsmasq/hosts --conf-dir=/home/dnsmasq/confs


# run whisper
cd /root/go/src/whisper # 因裡面有相對位置
go run main.go