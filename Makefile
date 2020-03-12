NAME=glu
ARCH=$(shell uname -m)
ORG=gliderlabs
VERSION=0.1.0

build:
	go build -o glu
	./glu container down
	./glu build linux,darwin
	rm ./glu
	docker build -t gliderlabs/glu:$(VERSION) .

test:
	docker run --rm gliderlabs/glu:$(VERSION) $(shell uname -s) | tar -xC /tmp
	/tmp/glu

install: build
	go install

deps:
	go get -d

release:
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
	docker tag gliderlabs/glu:$(VERSION) gliderlabs/glu:latest
	docker push gliderlabs/glu:latest
	docker push gliderlabs/glu:$(VERSION)

clean:
	rm -rf build release

.PHONY: build release
