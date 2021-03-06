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

tests:
		$(Q)echo "unit testing...."
		$(Q)go test ./pkg/...
		$(Q)go test ./test

verify: lint tests

prepare: fmt verify changelog
		$(Q)go mod tidy

clean:
		$(Q)rm -rf build

fmt:
		$(Q)echo "fixing imports and format...."
		$(Q)command -v goimports || GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
		$(Q)goimports -w .

mocks:
		$(Q)go get github.com/golang/mock/gomock
		$(Q)command -v mockgen || GO111MODULE=off go get github.com/golang/mock/mockgen
		$(Q)go generate ./...

changelog:
		$(Q)GO111MODULE=off go get -u github.com/typusomega/semantic-changelog-gen
		$(Q)echo "generating changelog...."
		$(Q)semantic-changelog-gen generate

update-golden:
		$(Q)go test ./test -update

build:
		$(Q)$(GOARGS) go build -gcflags "all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}" -o ./artifacts/semantic-changelog-gen ./cmd/semantic-changelog-gen/main.go
