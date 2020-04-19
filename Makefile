build:
	go build
	chmod u+x jcli-account-plugin

copy:
	cp jcli-account-plugin ~/.jenkins-cli/plugins

test:
	go test ./...
