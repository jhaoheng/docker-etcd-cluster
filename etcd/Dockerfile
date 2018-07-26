FROM        ubuntu:16.04

ENV         ETCD_PKG=https://github.com/coreos/etcd/releases/download/v3.3.9/etcd-v3.3.9-linux-amd64.tar.gz

RUN         apt-get update && apt-get install -y curl wget iputils-ping net-tools dnsutils

RUN         cd home && \
            curl -L $ETCD_PKG -o etcd.tar.gz && \
            tar xzvf etcd.tar.gz && \
            rm etcd.tar.gz && \
            cp etcd*/etcd /usr/local/bin/ && \
            cp etcd*/etcdctl /usr/local/bin/ && \
            rm -rf etcd*

RUN         etcd --version

COPY        ./script /home/etcd-script
RUN         chmod -R +x /home/etcd-script

