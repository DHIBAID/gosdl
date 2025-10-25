#!/bin/bash

# Install dependencies for Ebiten on Linux (Ubuntu/Debian)
sudo apt install libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config
go run github.com/hajimehoshi/ebiten/v2/examples/rotate@latest
go mod init gosdl
go mod tidy
