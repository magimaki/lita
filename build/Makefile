BINARY_NAME=lita
build:
	# build cli for local, use docker file to deployment

	# go build -ldflags="-s -w" -o $(BINARY_NAME) ../cmd/main.go
	go build -ldflags="-s -w" -o $(BINARY_NAME) ../cmd/main.go && upx -9 lita