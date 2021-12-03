setup:
	go mod tidy
	go install github.com/cosmtrek/air@latest
	go install mvdan.cc/gofumpt@latest
	
start:
	air

fmt:
	gofumpt -l -w .
