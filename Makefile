VERSION=`git describe --tags`
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.Version=${V} -X main.BuildTime=${BUILD_TIME}"
NAME=address

gen-code:
	protoc --go_out=. --go-grpc_out=. mygrpc/proto/*.proto


run: build
	./bin/$(NAME) -service grpc

build: clear
	go build ${LDFLAGS} -o ./bin/$(NAME) ./main.go
	./bin/$(NAME) -v

clear:
	rm -rf ./bin/$(NAME)

build-docker-img:
	docker build -t $(NAME):dev .
	docker rmi -f $$(docker images --filter “dangling=true” -q --no-trunc)