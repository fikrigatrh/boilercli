package main

//go:generate go run github.com/swaggo/swag/cmd/swag init
//go:generate go run github.com/google/wire/cmd/wire

func main() {
	server := ServerApp()
	server.Serve()
}
