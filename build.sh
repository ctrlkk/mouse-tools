#!/bin/bash

PROJECT_NAME="mouse-tools"

GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 wails build -platform windows/amd64 -o "${PROJECT_NAME}_windows_amd64.exe"
GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CGO_ENABLED=1 wails build -platform windows/386 -o "${PROJECT_NAME}_windows_386.exe"

wails build -tags webkit2_41 -platform linux/amd64 -o "${PROJECT_NAME}_linux_amd64"
