#FROM golang
FROM buildpack-deps:wheezy-scm

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

ENV GOLANG_VERSION 1.9.2
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 de874549d9a8d8d8062be05808509c09a88a248e77ec14eb77453530829ac02b
ENV GOPATH=/go
ENV GOSRC=$GOPATH/src
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# gcc for cgo
RUN apt-get update && \
    apt-get install -y --no-install-recommends make git bzip2 && \
	rm -rf /var/lib/apt/lists/* && \
    curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz && \
	echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - && \
	tar -C /usr/local -xzf golang.tar.gz && \
	rm golang.tar.gz && \
    mkdir -p "$GOPATH/src" "$GOPATH/bin" "$GOSRC/github.com/bayugyug" && chmod -R 777 "$GOPATH" && \
    cd $GOSRC/github.com/bayugyug && \
    git clone https://github.com/bayugyug/mongers-uagents.git && \
    cd $GOSRC/github.com/bayugyug/mongers-uagents && \
    go get -v && \
    go build -tags netgo -installsuffix netgo -installsuffix cgo -v -ldflags "-s -w -X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" . && \
    cd $GOSRC && \
    go install github.com/bayugyug/mongers-uagents && \
    cd $GOSRC/github.com/bayugyug/mongers-uagents  && \
    cp -f Docker/cleanup.sh $GOPATH/bin/ && \
    cp -f Docker/wheezy/go-wrapper /usr/local/bin/ && \
    chmod +x $GOPATH/bin/* && \
    $GOPATH/bin/cleanup.sh

#main entry
ENTRYPOINT [ "mongers-uagents" ]
