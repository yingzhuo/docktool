TIMESTAMP  	:= $(shell /bin/date "+%F %T")
VERSION		:= latest
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

release-dryrun:
	BuildGitBranch=$(shell git describe --all) \
	 	BuildGitRev=$(shell git rev-list --count HEAD) \
	 	BuildGitCommit=$(shell git rev-parse HEAD) \
	 	goreleaser release --skip-publish --snapshot --rm-dist

release:
	BuildGitBranch=$(shell git describe --all) \
	 	BuildGitRev=$(shell git rev-list --count HEAD) \
	 	BuildGitCommit=$(shell git rev-parse HEAD) \
		goreleaser release --rm-dist

install:
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 sudo go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o /usr/local/bin/$(NAME)
	@sudo chmod a+x /usr/local/bin/$(NAME)

uninstall:
	@sudo rm -rf /usr/local/bin/$(NAME) &> /dev/null

.PHONY: fmt clean release-dryrun release install uninstall
