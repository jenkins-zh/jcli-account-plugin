build: pre-build
	go build
	chmod u+x jcli-account-plugin

vet:
	go vet ./...

pre-build: fmt vet
	export GO111MODULE=on
	export GOPROXY=https://goproxy.io
	go mod tidy

copy:
	cp jcli-account-plugin ~/.jenkins-cli/plugins

test:
	go test ./...

fmt:
	go fmt .
	gofmt -s -w .