all: build install

build:
	go build -o пше

install:
	rm ~/bin/пше
	mv пше ~/bin/пше
	chmod +x ~/bin/пше