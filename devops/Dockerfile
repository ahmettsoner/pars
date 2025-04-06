FROM ubuntu:22.04

# Gerekli sistem araçları
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    make \
    cmake \
    build-essential \
    unzip \
    ca-certificates \
    gnupg \
    lsb-release \
    software-properties-common

# Node.js 22 kurulumu
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g npm

# Go 1.22 kurulumu
RUN wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz && \
    rm go1.22.0.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Node modülleri kurulumu ve build
RUN npm install -g git-release-manager

RUN useradd -m builder
USER builder
WORKDIR /workspace

