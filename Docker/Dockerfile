#start
FROM golang:1.8

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

#environment 
ENV GOHOME=
ENV GOPATH=$GOHOME/go
ENV GOSRC=$GOPATH/src
ENV PATH=$PATH:$GOPATH/bin

#prepare
RUN apt-get install -y --no-install-recommends git && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p $GOPATH/bin  $GOSRC/github.com/bayugyug && \
    cd $GOSRC/github.com/bayugyug && \
    git clone https://github.com/bayugyug/mongers-uagents.git && \
    cd $GOSRC/github.com/bayugyug/mongers-uagents && \
    go get -v && go build -tags netgo -installsuffix netgo -installsuffix cgo -v -ldflags "-s -w -X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" . && \
    cd $GOSRC && go install github.com/bayugyug/mongers-uagents && \
    cd $GOSRC/github.com/bayugyug/mongers-uagents && \ 
    chmod +x $GOPATH/bin/*

#main entry
ENTRYPOINT ["mongers-uagents"]
