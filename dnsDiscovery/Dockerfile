FROM      ubuntu:16.04

# dns
RUN         apt-get update && apt-get install dnsmasq vim dnsutils wget iputils-ping curl net-tools -y
RUN         dnsmasq --version
RUN         mkdir -p /home/dnsmasq/hosts && \
            mkdir -p /home/dnsmasq/confs

# golang
RUN         cd /usr/local && \
            wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz && \
            tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz && \
            rm go1.10.3.linux-amd64.tar.gz
          
RUN         ln -s /usr/local/go/bin/* /bin && \
            mkdir -p /root/go/src && \
            mkdir -p /root/go/bin && \
            mkdir -p /root/go/pkg

COPY        ./whisper /root/go/src/whisper
RUN         chmod -R +x /root/go/src/whisper
EXPOSE      8080/tcp

COPY        ./entrypoint.sh /
RUN         chmod +x /entrypoint.sh
ENTRYPOINT  /entrypoint.sh