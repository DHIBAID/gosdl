debug:
	clear
	__NV_PRIME_RENDER_OFFLOAD=1 __GLX_VENDOR_LIBRARY_NAME=nvidia SDL_VIDEO_DRIVER=x11 SDL_RENDER_DRIVER=opengl go run main.go

watch:
	air

check:
	go vet ./...
	go fmt ./...
	golangci-lint run

build:
	__NV_PRIME_RENDER_OFFLOAD=1 __GLX_VENDOR_LIBRARY_NAME=nvidia SDL_VIDEO_DRIVER=x11 SDL_RENDER_DRIVER=opengl go build -o bin/app main.go