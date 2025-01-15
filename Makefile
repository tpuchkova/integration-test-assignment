all:
	go run cmd/main.go

lint:
	gofumpt -l -w .
	gci write --skip-generated -s standard,default .
	golangci-lint run --max-same-issues=0 --max-issues-per-linter=0