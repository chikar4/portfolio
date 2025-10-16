run:
	@go run backend/cmd/main.go

build:
	@go build -o backend/bin/portfolio backend/cmd/main.go

clean:
	@rm -rf backend/bin/* 
