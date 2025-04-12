FROM ubuntu:22.04

# Gerekli sistem araçları
RUN apt-get update 
RUN apt-get install -y curl
RUN apt-get install -y wget
RUN apt-get install -y git
RUN apt-get install -y make
RUN apt-get install -y cmake
RUN apt-get install -y build-essential

# Node.js 22 kurulumu
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g npm


# Go 1.22 kurulumu
RUN wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz && \
    rm go1.22.0.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"


RUN apt-get install -y zip
RUN apt-get install -y unzip
RUN apt-get install -y p7zip-full
RUN apt-get install -y rar
RUN apt-get install -y ca-certificates
RUN apt-get install -y gnupg
RUN apt-get install -y lsb-release
RUN apt-get install -y software-properties-common
RUN apt-get install -y pandoc
RUN apt-get install -y devscripts
RUN apt-get install -y dh-make
RUN apt-get install -y dh-golang
RUN apt-get install -y debhelper
RUN apt-get install -y lintian
RUN apt-get install -y fakeroot
RUN apt-get install -y golang-any
RUN apt-get install -y rpm

# Node modülleri kurulumu ve build
RUN npm install -g git-release-manager@0.0.12

RUN useradd -m builder
USER builder
WORKDIR /workspace

