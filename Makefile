all:
	go build -o arbiter app/arbiter/main.go
	go build -o arbiter_rpc app/rpc/main.go
	go build -o arbiter_web app/web/main.go
	go build -o arbiter_web app/email/main.go

linux:
	GOARCH=amd64 GOOS=linux go build -o arbiter app/arbiter/main.go
	GOARCH=amd64 GOOS=linux go build -o arbiter_rpc app/rpc/main.go
	GOARCH=amd64 GOOS=linux go build -o arbiter_web app/web/main.go
	GOARCH=amd64 GOOS=linux go build -o arbiter_email app/email/main.go