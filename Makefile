all: build

build :
	go get -v
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix netgo -installsuffix cgo -v -ldflags "-s -w -X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" .

test : build
	go test *.go > testrun.txt
	golint > lint.txt
	go tool vet -v . > vet.txt
	gocov test github.com/bayugyug/mongers-uagents | gocov-xml > coverage.xml
	go test *.go -bench=. -test.benchmem -v 2>/dev/null | gobench2plot > benchmarks.xml

testrun : clean test
	time go test -v -bench=. -benchmem -dummy >> testrun.txt 2>&1

prepare : build
	cp mongers-uagents Docker/mongers-uagents

docker-devel : prepare
	-@sudo docker rmi -f bayugyug/mongers-uagents 2>/dev/null || true
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/mongers-uagents .

docker-wheezy: prepare
	-@sudo docker rmi -f bayugyug/mongers-uagents 2>/dev/null || true
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/mongers-uagents -f  wheezy/Dockerfile .

docker-scratch: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/mongers-uagents:scratch -f  scratch/Dockerfile .

docker-latest: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/mongers-uagents -f  scratch/Dockerfile .

docker-alpine: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/mongers-uagents:alpine  -f  alpine/Dockerfile .

clean:
	rm -f mongers-uagents Docker/mongers-uagents
	rm -f benchmarks.xml coverage.xml vet.txt lint.txt testrun.txt

re: clean all

