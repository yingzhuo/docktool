TIMESTAMP  	:= $(shell /bin/date "+%F %T")
VERSION		:= v1.2.0
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
	@rm -rf $(CURDIR)/_dist &> /dev/null

release: clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_dist/$(NAME)
	cp $(CURDIR)/Dockerfile $(CURDIR)/_dist
	docker image build -t yingzhuo/$(NAME):$(VERSION) $(CURDIR)/_dist
	docker image tag yingzhuo/$(NAME):$(VERSION) yingzhuo/$(NAME):latest
	docker push yingzhuo/$(NAME):$(VERSION)
	docker push yingzhuo/$(NAME):latest

install:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 sudo go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o /usr/local/bin/$(NAME)
	@sudo chmod a+x /usr/local/bin/$(NAME)

uninstall:
	@sudo rm -rf /usr/local/bin/$(NAME) &> /dev/null

.PHONY: fmt clean release install uninstall