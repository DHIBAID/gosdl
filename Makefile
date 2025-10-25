debug:
	clear
	__NV_PRIME_RENDER_OFFLOAD=1 __GLX_VENDOR_LIBRARY_NAME=nvidia SDL_VIDEO_DRIVER=x11 SDL_RENDER_DRIVER=opengl go run main.go

linter:
	go vet ./...
	go fmt ./...
	golangci-lint run

build:
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -ldflags="-s -w" -trimpath -o bin/app main.go
