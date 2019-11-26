TIMESTAMP  	:= $(shell /bin/date "+%F %T")
VERSION		:= 1.0.10
NAME		:= docktool
LDFLAGS		:= -s -w \
			   -X 'main.BuildVersion=$(VERSION)' \
			   -X 'main.BuildGitBranch=$(shell git describe --all)' \
			   -X 'main.BuildGitRev=$(shell git rev-list --count HEAD)' \
			   -X 'main.BuildGitCommit=$(shell git rev-parse HEAD)' \
			   -X 'main.BuildDate=$(shell /bin/date "+%F %T +08:00")'

fmt:
	@go fmt ./...
	@go mod tidy

clean:
	@rm -rf $(CURDIR)/_bin &> /dev/null

build: clean
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-linux-amd64-$(VERSION)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-darwin-amd64-$(VERSION)

install: uninstall build
	sudo cp $(CURDIR)/_bin/$(NAME)-darwin-amd64-$(VERSION) /usr/local/bin/$(NAME)
	@sudo chmod a+x /usr/local/bin/$(NAME)

uninstall:
	@sudo rm -rf /usr/local/bin/$(NAME)-* &> /dev/null
	@sudo rm -rf /usr/local/bin/frep &> /dev/null

github: fmt
	@git add . &> /dev/null
	@git commit -m "$(TIMESTAMP)"
	@git push origin master

.PHONY: fmt clean build install uninstall github
