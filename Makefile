TIMESTAMP  	:= $(shell /bin/date "+%F %T")
VERSION		:= v1.1.1
NAME		:= docktool
LDFLAGS		:= -s -w \
			   -X 'main.BuildVersion=$(VERSION)' \
			   -X 'main.BuildGitBranch=$(shell git describe --all)' \
			   -X 'main.BuildGitRev=$(shell git rev-list --count HEAD)' \
			   -X 'main.BuildGitCommit=$(shell git rev-parse HEAD)' \
			   -X 'main.BuildDate=$(shell /bin/date "+%F %T +08:00")' \
			   -X 'main.BuildBy=make'

fmt:
	@go fmt ./...
	@go mod tidy

clean:
	@rm -rf $(CURDIR)/_bin &> /dev/null

build: clean
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-linux-amd64-$(VERSION)
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-darwin-amd64-$(VERSION)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-windows-amd64-$(VERSION).exe

install:
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 sudo go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o /usr/local/bin/$(NAME)
	@sudo chmod a+x /usr/local/bin/$(NAME)

uninstall:
	@sudo rm -rf /usr/local/bin/$(NAME) &> /dev/null

.PHONY: fmt clean build install uninstall
