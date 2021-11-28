export tag=v2.0
root:
	export ROOT=github.com/zhdd99/CloudNative/homework_module8

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t zhdd99/httpserver:${tag} .

push: release
	echo "pushing zhdd99/httpserver"
	docker push zhdd99/httpserver:${tag}
