default: build-all

.PHONY: build-all
build-all: linux-intel mac-intel windows-intel

.PHONY: mac-intel
mac-intel:
	env GOOS=darwin GOARCH=amd64 go build -o build/zfix
	chmod u+x build/zfix

.PHONY: linux-intel
linux-intel:
	env GOOS=linux GOARCH=amd64 go build -o build/zfix-nix
	chmod u+x build/zfix-nix

.PHONY: windows-intel
windows-intel:
	env GOOS=windows GOARCH=amd64 go build -o build/zfix.exe
