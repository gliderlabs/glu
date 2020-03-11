.PHONY: build

build:
	go get -d
	go build -o glu
	./glu container down
	./glu build linux,darwin
	rm ./glu
	docker build -t gliderlabs/glu .

install:
	go install
