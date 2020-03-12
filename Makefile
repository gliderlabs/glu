glu
ARCH=$(shell uname -m)
ORG=gliderlabs
VERSION=0.1.0

build:
	go build -o glu
	./glu container down
	./glu build linux,darwin
	rm ./glu
	docker build -t gliderlabs/glu .

test:
	docker run --rm gliderlabs/glu $(shell uname -s) | tar -xC /tmp
	/tmp/glu

install: build
	go install

deps:
	go get -d

release:
  docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
  docker push gliderlabs/glu

clean:
	rm -rf build release

.PHONY: build release
