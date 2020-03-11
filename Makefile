.PHONY: build

build:
	go get -d
	ls -lah
	go build
	ls -lah
	./glu container down
	./glu build linux,darwin
	rm ./glu
	docker build -t gliderlabs/glu .

install:
	go install
