all: build/i-wanna-play-minecraft build/linux/amd64/i-wanna-play-minecraft build/linux/arm64/i-wanna-play-minecraft build/darwin/amd64/i-wanna-play-minecraft build/darwin/arm64/i-wanna-play-minecraft build/windows/amd64/i-wanna-play-minecraft.exe build/windows/arm64/i-wanna-play-minecraft.exe

clean:
	rm -rf build

build/i-wanna-play-minecraft:
	go build -o build/i-wanna-play-minecraft cmd/main.go

build/linux/amd64/i-wanna-play-minecraft:
	GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/i-wanna-play-minecraft cmd/main.go

build/linux/arm64/i-wanna-play-minecraft:
	GOOS=linux GOARCH=arm64 go build -o build/linux/arm64/i-wanna-play-minecraft cmd/main.go

build/darwin/amd64/i-wanna-play-minecraft:
	GOOS=darwin GOARCH=amd64 go build -o build/darwin/amd64/i-wanna-play-minecraft cmd/main.go

build/darwin/arm64/i-wanna-play-minecraft:
	GOOS=darwin GOARCH=arm64 go build -o build/darwin/arm64/i-wanna-play-minecraft cmd/main.go

build/windows/amd64/i-wanna-play-minecraft.exe:
	GOOS=windows GOARCH=amd64 go build -o build/windows/amd64/i-wanna-play-minecraft.exe cmd/main.go

build/windows/arm64/i-wanna-play-minecraft.exe:
	GOOS=windows GOARCH=arm64 go build -o build/windows/arm64/i-wanna-play-minecraft.exe cmd/main.go