all: recompile

build:
	go build -o build/go-dc-bot
run:
	go run .
clean:
	rm -rf build/*
compile:
	GOOS=windows GOARCH=amd64 go build -o build/go-dc-bot-windows-x64.exe
	GOOS=windows GOARCH=386 go build -o build/go-dc-bot-windows-x32.exe
	GOOS=linux GOARCH=386 go build -o build/go-dc-bot-linux-i386
	GOOS=linux GOARCH=amd64 go build -o build/go-dc-bot-linux-amd64
	GOOS=linux GOARCH=arm go build -o build/go-dc-bot-linux-arm32
	GOOS=linux GOARCH=arm64 go build -o build/go-dc-bot-linux-arm64
	GOOS=freebsd GOARCH=386 go build -o build/go-dc-bot-freebsd-386

rebuild: clean build
recompile: clean compile