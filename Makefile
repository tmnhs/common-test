# Go makefile

#export env
#basic information
ProjectName := "common-test"

PROJECTBASE 	:= $(shell pwd)
PROJECTBIN 	:= $(PROJECTBASE)/bin
TIMESTAMP   := $(shell /bin/date "+%F %T")

#change to deploy environment
MainFile := ./cmd/main.go

#compile ldflags
LDFLAGS		:= -s -w \
			   -X 'main.BuildGitBranch=$(shell git describe --all)' \
			   -X 'main.BuildGitRev=$(shell git rev-list --count HEAD)' \
			   -X 'main.BuildGitCommit=$(shell git rev-parse HEAD)' \
			   -X 'main.BuildDate=$(shell /bin/date "+%F %T")'


linux-dev: clean
	@echo "install linux amd64 dev version"
	@go mod tidy
	@echo "building project "$(ProjectName)"..."
	@CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -v -o $(PROJECTBIN)/$(ProjectName) $(MainFile)
	@chmod +x $(PROJECTBIN)/
	@echo "build success."

local-dev: clean
	@echo "install local dev version"
	@go mod tidy

	@echo "building project "$(ProjectName)"..."
	@CGO_ENABLED=0 go build -v  -o $(PROJECTBIN)/$(ProjectName) $(MainFile)

	@chmod +x $(PROJECTBIN)/
	@echo "build success."

gitpush: clean fmt
	git add .
	git commit -m "$(m) changed at $(TIMESTAMP)"
	git push
fmt:
	@go fmt $(PROJECTBASE)/...
	@echo "hello"
	@go mod tidy

clean:
	@#echo $(PROJECTBIN)
	@rm -rf $(PROJECTBIN)/* &>/dev/null
depend:
	go mod download
gitpull: fmt
	git add .
	git commit -m "$(m) changed at $(TIMESTAMP)"
	git pull
.PHONY: fmt clean git