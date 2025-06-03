# BoilerCLI Enhanced Interactive

🚀 A powerful CLI tool for generating Go microservice boilerplates with interactive prompts and embedded templates.

## Features

- 📦 Interactive project generation
- 🌐 Support for Gin and Fiber web frameworks
- 📁 Embedded template system
- 🎯 Clean architecture structure
- ⚡ Ready-to-use configurations
- 🔧 Customizable module paths

## Project Structure

```
boilercli-enhanced-interactive/
├── cmd/
│   ├── create.go      # Interactive creation command
│   └── root.go        # Root cobra command
├── template/
│   ├── fiber/         # Fiber framework templates
│   └── gin/           # Gin framework templates
├── embed.go           # Embedded file system
├── go.mod
├── main.go            # Main entry point
└── README.md
```

## Installation

### Build from Source

```bash
# Clone the repository
git clone <repo-url>
cd boilercli-enhanced-interactive

# Build the binary
go build -o boilercli .

# Make it executable (Linux/macOS)
chmod +x boilercli

# Optional: Move to PATH
sudo mv boilercli /usr/local/bin/
```

## Usage

### Interactive Mode

Run the CLI tool and follow the interactive prompts:

```bash
./boilercli create
```

The tool will prompt you for:
- 📦 **Service/Project Name**: Your project name
- 📁 **Output Directory**: Where to create the project (e.g., `./services`)
- 📦 **Go Module Path**: Your Go module path (e.g., `github.com/myorg/myservice`)
- 🌐 **HTTP Framework**: Choose between `gin` or `fiber`

### Example Session

```bash
$ ./boilercli create

╔════════════════════════════════════╗
║ 🚀 BoilerCLI - Project Generator  ║
╚════════════════════════════════════╝

📦 Enter service/project name: my-awesome-api
📁 Enter output directory (e.g., ./services): ./projects
📦 Enter Go module path (e.g., github.com/myorg/myservice): github.com/myorg/my-awesome-api
🌐 Enter HTTP Router / Web Frameworks (gin or fiber): gin

✅ Project created at ./projects/my-awesome-api
Your Project Generated in: ./projects/my-awesome-api
```

## Generated Project Structure

The generated project follows clean architecture principles:

```
your-project/
├── configs/
│   ├── infra/          # Database, Redis configurations
│   ├── logger/         # Logger setup
│   └── router/         # Router configuration
├── external/           # External service integrations
├── internal/
│   ├── dto/           # Data Transfer Objects
│   ├── error/         # Error definitions
│   ├── handler/       # HTTP handlers
│   ├── middleware/    # Custom middleware
│   ├── model/         # Data models
│   ├── repository/    # Data access layer
│   ├── usecase/       # Business logic
│   ├── transport/     # Transport layer
│   └── utils/         # Internal utilities
├── transport/         # Transport configurations
├── utils/            # Global utilities
├── env_dev.yml.example # Configuration example
├── go.mod            # Go module file
└── README.md         # Project documentation
```

## Development

### Prerequisites

- Go 1.19 or higher
- Git

### Setup Development Environment

```bash
# Clone the repository
git clone <your-repo-url>
cd boilercli-enhanced-interactive

# Install dependencies
go mod tidy

# Run the application
go run main.go create
```

### Building

```bash
# Build for current platform
go build -o boilercli .

# Build for multiple platforms
# Linux
GOOS=linux GOARCH=amd64 go build -o boilercli-linux .

# Windows
GOOS=windows GOARCH=amd64 go build -o boilercli.exe .

# macOS
GOOS=darwin GOARCH=amd64 go build -o boilercli-macos .
```

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Generated Project Quick Start

After generating a project, navigate to it and:

```bash
cd your-generated-project

# Install dependencies
go mod tidy

# Copy configuration file
cp env_dev.yml.example env/env_dev.yml

# Edit configuration as needed
nano env/env_dev.yml

# Generate Swagger documentation (if applicable)
go run github.com/swaggo/swag/cmd/swag init

# Generate Wire dependency injection (if applicable)
go run github.com/google/wire/cmd/wire

# Run the application
go run main.go
```

## Configuration

The generated projects use Viper for configuration management and support:

- **YAML files**: `env_dev.yml`, `env_prod.yml`
- **Environment variables**: Prefixed with service name
- **Command line flags**: Standard Go flags

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Supported Frameworks

- **Gin**: Fast HTTP web framework
- **Fiber**: Express-inspired web framework

## Dependencies

- `github.com/spf13/cobra`: CLI framework
- `text/template`: Go template engine
- `embed`: Go embedded file system

## Acknowledgments

- Created by **gatgat**
- Assisted by **Claude AI**, **Gemini**, and **ChatGPT**
- Built with ❤️ for the Go community

## Support

If you encounter any issues or have questions:

1. Check the [Issues](../../issues) section
2. Create a new issue with detailed information
3. Provide steps to reproduce any bugs

---

**Happy coding! 🚀**
