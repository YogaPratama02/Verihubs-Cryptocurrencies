Arguments := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

setup:
	go mod tidy

env:
	cp .env.example .env

serve:
	go run main.go
