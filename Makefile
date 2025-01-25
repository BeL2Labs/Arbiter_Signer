all:
	go build -o arbiter app/arbiter/main.go
	go build -o keystore-generator app/keystore-generator/main.go

linux:
	GOARCH=amd64 GOOS=linux go build -o arbiter app/arbiter/main.go
	GOARCH=amd64 GOOS=linux go build -o keystore-generator app/keystore-generator/main.go