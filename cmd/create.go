package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"text/template"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Interactively create a new service from template",
	Run: func(cmd *cobra.Command, args []string) {
		printLogo()

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("📦 Enter service/project name: ")
		projectName, _ := reader.ReadString('\n')
		projectName = strings.TrimSpace(projectName)

		fmt.Print("📁 Enter output directory (e.g., ./services): ")
		outputPath, _ := reader.ReadString('\n')
		outputPath = strings.TrimSpace(outputPath)

		fmt.Print("📦 Enter Go module path (e.g., github.com/myorg/myservice): ")
		modulePath, _ := reader.ReadString('\n')
		modulePath = strings.TrimSpace(modulePath)

		fmt.Print("🌐 Enter HTTP Router / Web Frameworks (gin or fiber): ")
		router, _ := reader.ReadString('\n')
		router = strings.TrimSpace(router)
		src := filepath.Join(fmt.Sprintf("template/%s", router))
		dst := filepath.Join(outputPath, projectName)

		err := copyAndReplace(src, dst, "boilerplate", modulePath)
		if err != nil {
			fmt.Println("❌ Error creating project:", err)
			return
		}
		fmt.Println("✅ Project created at", dst)

		err = initializeTemplates(fmt.Sprintf("%s/%s", outputPath, projectName), TemplateData{
			PROJECT_NAME: projectName,
			MODULE_NAME:  modulePath,
			SERVICE_NAME: projectName,
		})
		if err != nil {
			fmt.Println("❌ Error creating:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func copyAndReplace(src, dst, oldImport, newImport string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		target := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}

		input, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		content := strings.ReplaceAll(string(input), oldImport, newImport)
		return os.WriteFile(target, []byte(content), info.Mode())
	})
}

func printLogo() {
	fmt.Println(`
╔════════════════════════════════════╗
║ 🚀 BoilerCLI - Project Generator  ║
╚════════════════════════════════════╝
`)
}

// TemplateData holds the values to replace in the templates
type TemplateData struct {
	PROJECT_NAME string
	MODULE_NAME  string
	SERVICE_NAME string
}

func initializeTemplates(templateDir string, templateData TemplateData) error {
	// Template files with placeholders
	templateFiles := map[string]string{
		"go.mod": `module {{.MODULE_NAME}}

go 1.23

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/spf13/viper v1.16.0
	gorm.io/gorm v1.25.4
	gorm.io/driver/postgres v1.5.2
	github.com/redis/go-redis/v9 v9.0.5
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/fx v1.20.0
)
`,
		"env_dev.yml.example": `
envLib:
  app:
  	name: "{{.PROJECT_NAME}}"
    host: "0.0.0.0"
    port: 9898
    version: "v1.0.0"
    envPrefix: "YOUR_SERVICE"

appEnvMode:
  mode: "dev"
  testPathPrefix: "../../../"
  debugMode: false
  ginMode: "debug"
  isPrettyLog: false

databaseConfig:
  dbPostgres:
    main_db:
      name: "{{.PROJECT_NAME}}"
      host: "localhost"
      port: "5432"
      user: "user"
      pass: "password"
      tz: "Asia/Jakarta"
    test_db:
      name: "{{.PROJECT_NAME}}"
      host: "localhost"
      port: "5432"
      user: "user"
      pass: "password"
      tz: "Asia/Jakarta"
  dbMysql:
    main_db:
      name: "{{.PROJECT_NAME}}"
      host: "localhost"
      port: "3306"
      user: "user"
      pass: "password"
      tz: "Asia%2FJakarta"

logConfig:
  elastic:
    host: ""
    port: ""
    index: ""
    user: ""
    pass: ""
  hookElasticEnabled: false
  gspaceChat:
    isEnabled: ""
    space_id: ""
    space_secret: ""
    space_token: ""
    serviceName: ""

external:
  serviceName:
	baseUrl: ""

redis:
  enable: true
  host: "localhost"
  port: "1111"
  user: "user"
  password: "password"
  db: 0
  usingTLS: "true"
  tlsCACert: "path_location/ca.pem"
  tlsCert: "path_location/user.crt"
  tlsKey: "path_location/user.key"
  poolSize: 5
  minIdleConn: 5

request:

server:
  shutdown:
    cleanup_period_seconds: "3"
    grace_period_seconds: "4"
  timeout:
    duration: "10s"

cors:
  allowOrigins: "*,http://localhost:8082"
  allowMethods: "GET,POST,PUT,PATCH,DELETE,OPTION"
  allowHeaders: "Origin, Content-Type, Accept, Authorization"
  allowCredentials: true
  exposeHeaders: "Content-Length"
`,
		"README.md": `# {{.PROJECT_NAME}}

This project was generated using the Go boilerplate CLI tool.

## Project Structure

- **configs/**: Configuration files and providers
  - **infra/**: Infrastructure configurations (database, redis)
  - **logger/**: Logger configuration
  - **router/**: Router configuration
- **external/**: External service integrations
- **internal/**: Internal application code
  - **dto/**: Data Transfer Objects
  - **error/**: Error definitions
  - **handler/**: HTTP handlers
  - **middleware/**: Custom middleware
  - **model/**: Data models
  - **repository/**: Data access layer
  - **usecase/**: Business logic layer
  - **transport/**: Transport layer
  - **utils/**: Internal utilities
- **transport/**: Transport related code
- **utils/**: Global utilities

## Getting Started

1. Install dependencies:
   ` + "```bash" + `
   go mod tidy
   ` + "```" + `

2. Generate swagger:
   ` + "```bash" + `
   go run github.com/swaggo/swag/cmd/swag init
   ` + "```" + `

3. Generate wire:
   ` + "```bash" + `
   go run github.com/google/wire/cmd/wire
   ` + "```" + `

4. Copy and modify the configuration:
	` + "```bash" + `
	# Development
	cp env_dev.yml.example env/env_dev.yml

	# Production
	# Option 1: Rename the dev file
	mv env_dev.yml env/env_prod.yml

	# Option 2: Create a new prod file
	touch env/env_prod.yml
	` + "```" + `

5. Run the application:
   ` + "```bash" + `
   go run main.go
   ` + "```" + `

## Configuration

The application uses Viper for configuration management. Configuration can be provided via:
- YAML files (env_dev.yaml)
- Environment variables
- Command line flags

## Development

This project follows a clean architecture pattern with clear separation of concerns:

- **Transport Layer**: Handles HTTP requests/responses
- **Handler Layer**: Processes HTTP requests
- **Usecase Layer**: Contains business logic
- **Repository Layer**: Handles data persistence
- **Model Layer**: Defines data structures

Happy coding! 🚀

Created by gatgat and assisted by Claude ai, Gemini and chatGPT
`,
		"Makefile": `# Build the application
build:
	go build -o bin/{{.PROJECT_NAME}} main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Install dependencies
deps:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	rm -rf bin/

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Run the application in development mode
dev:
	air

.PHONY: build run test test-coverage deps clean fmt lint dev
`,
	}

	for filename, content := range templateFiles {
		fullPath := filepath.Join(templateDir, filename)

		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory for template %s: %v", filename, err)
		}

		// Parse and execute the template
		tmpl, err := template.New(filename).Parse(content)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %v", filename, err)
		}

		file, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", filename, err)
		}
		defer file.Close()

		if err := tmpl.Execute(file, templateData); err != nil {
			return fmt.Errorf("failed to execute template for file %s: %v", filename, err)
		}
	}

	fmt.Printf("\nYour Project Generated in: %s\n", templateDir)
	return nil
}
