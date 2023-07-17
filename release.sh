#!/bin/zsh
rm -rf target && mkdir target

# build for windows-x86_64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o target/pg-windows-amd64.exe
# build for macOS-x86_64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target/pg-darwin-amd64
#build for macOS-arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o target/pg-darwin-arm64
# build for linux-x86_64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/pg-linux-amd64

# calculate sha256 checksum
cd target
echo "sha256 checksum" > checksum
sha256sum pg-windows-amd64.exe >> checksum
sha256sum pg-darwin-amd64 >> checksum
sha256sum pg-darwin-arm64 >> checksum
sha256sum pg-linux-amd64 >> checksum
