GOOS ?= windows


all: dist/ui dist/runner dist/setup dist/data dist/log

dist:
	mkdir dist

dist/data: dist
	mkdir -p dist/data

dist/log: dist
	mkdir -p dist/log

dist/ui: dist ui/dist
	cp -r ui/dist dist/ui

ui/dist:
	cd ui && yarn build

dist/runner: dist
ifeq ($(GOOS),windows)
	go build -ldflags -H=windowsgui -o dist/runner.exe ./cmd/runner/main.go
else
	go build -o dist/runner ./cmd/runner/main.go
endif

dist/setup: dist
ifeq ($(GOOS),windows)
	go build -ldflags -H=windowsgui -o dist/setup.exe ./cmd/setup/main.go
else
	go build -o dist/setup ./cmd/setup/main.go
endif

clean:
	rm -rf dist ui/dist
.PHONY: clean