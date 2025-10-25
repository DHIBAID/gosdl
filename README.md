# GoSDL

GoSDL is a Go-based project leveraging SDL2 (Simple DirectMedia Layer) to create graphical applications. This project is designed to provide a robust framework for rendering, physics, and window management, making it easier to build interactive and visually appealing applications.

## Features

- **Window Management**: Configurable window settings for resolution, fullscreen, and more.
- **Rendering**: Efficient rendering pipeline with support for anti-aliasing and layout management.
- **Physics**: Basic physics configurations for interactive applications.
- **Font Support**: Includes font management for text rendering.
- **Modular Design**: Organized into distinct modules for better maintainability.

## Project Structure

```
.
├── go.mod          # Go module file
├── main.go         # Entry point of the application
├── Makefile        # Build and run commands
├── config/         # Configuration files
│   ├── physics.go  # Physics-related configurations
│   └── window.go   # Window-related configurations
├── fonts/          # Fonts directory
├── renderer/       # Rendering logic
│   ├── antialiasing.go
│   ├── draw.go
│   ├── init.go
│   ├── layout.go
│   └── update.go
├── setup/          # Setup scripts
│   └── setup.go    # Go script for setup
├── tmp/            # Temporary files
│   └── main        # Temporary build output
```

## Getting Started

### Prerequisites

- **Go**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org/).
- **SDL2**: Install SDL2 on your system. For Linux, you can use your package manager:

  ```bash
  sudo apt-get install libsdl2-dev
  ```

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/DHIBAID/gosdl.git
   ```

2. Navigate to the project directory:

   ```bash
   cd gosdl
   ```

3. Run the setup script:

   ```bash
   go run setup/setup.go
   ```

4. Install dependencies:

   ```bash
   go mod tidy
   ```

### Running the Application

To build and run the application, use the provided `Makefile`:

```bash
make build
./bin/app
```

Alternatively, you can build and run the application directly with:

```bash
go build -o bin/app main.go
./bin/app
```

### Makefile Commands

- **`make build`**: Builds the application and outputs the binary to `tmp/main`.
- **`make debug`**: Runs the application.
- **`make clean`**: Cleans up temporary files and build artifacts.
