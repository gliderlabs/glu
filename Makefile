.PHONY: build

build:
	go get || true
	go install # for us
	go build
	./glu container down
	./glu build linux,darwin
	rm ./glu
	docker build -t gliderlabs/glu .
