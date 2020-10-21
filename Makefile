VERSION = `git describe --abbrev=0 --tags`
GHRFLAGS =
.PHONY: build release

all: install

build:
	goxc -d=pkg -pv=v$(VERSION) -os="linux darwin windows"

release:
	ghr -delete -u temamagic  $(GHRFLAGS) $(VERSION) pkg/$(VERSION)

install:
	go build -o пше
	rm -f ~/bin/пше
	mv пше ~/bin/пше
	chmod +x ~/bin/пше

report:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

check:
	golangci-lint run ./...


changelog:
	git-chglog -o CHANGELOG.md