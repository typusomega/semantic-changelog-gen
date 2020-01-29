# kernel-style V=1 build verbosity
ifeq ("$(origin V)", "command line")
       BUILD_VERBOSE = $(V)
endif

ifeq ($(BUILD_VERBOSE),1)
       Q =
else
       Q = @
endif

export CGO_ENABLED:=0

all: verify build

lint:
		$(Q)echo "linting...."
		$(Q)command -v golangci-lint || GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
		$(Q)golangci-lint run -E gofmt -E golint -E goconst -E gocritic -E golint -E gosec -E maligned -E nakedret -E prealloc -E unconvert -E gocyclo -E scopelint -E goimports
		$(Q)echo linting OK

test:
		$(Q)echo "unit testing...."
		$(Q)go test ./pkg/...

verify: lint test

prepare: fmt verify
		$(Q)go mod tidy

clean:
		$(Q)rm -rf build

fmt:
		$(Q)echo "fixing imports and format...."
		$(Q)command -v goimports || GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
		$(Q)goimports -w .

mocks:
		$(Q)go get github.com/golang/mock/gomock
		$(Q)go install github.com/golang/mock/mockgen
		$(Q)go generate ./...

build:
		$(Q)$(GOARGS) go build -gcflags "all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}" -o ./artifacts/semantic-changelog-gen ./cmd/semantic-changelog-gen/main.go
