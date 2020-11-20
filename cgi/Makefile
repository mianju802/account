build:
		GOOS=linux GOARCH=amd64 go build
		docker build -t cgi .
run:
		docker run -e MICRO_REGISTRY=mdns cgi