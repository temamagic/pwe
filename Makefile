VERSION = `git describe --abbrev=0 --tags`
GHRFLAGS =
.PHONY: build release

all: install

build:
	goxc -d=pkg -pv=$(VERSION) -os="linux darwin windows"

release:
	ghr -delete -u temamagic  $(GHRFLAGS) $(VERSION) pkg/$(VERSION)

install:
	go build -o пше
	rm -f ~/bin/пше
	mv пше ~/bin/пше
	chmod +x ~/bin/пше