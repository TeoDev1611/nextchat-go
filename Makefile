setup:
	go mod tidy
	go install github.com/cosmtrek/air@latest

start:
	air
