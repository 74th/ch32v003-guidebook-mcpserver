#!/bin/bash
./generate_documents_blob.py
mkdir release
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o release/ch32v003-mcp-server_linux_amd64 .
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o release/ch32v003-mcp-server_linux_arm64 .
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o release/ch32v003-mcp-server_darwin_amd64 .
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o release/ch32v003-mcp-server_darwin_arm64 .
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o release/ch32v003-mcp-server_windows_amd64.exe .
GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -o release/ch32v003-mcp-server_windows_arm64.exe .