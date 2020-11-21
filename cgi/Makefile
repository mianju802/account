build:
		GOOS=linux CGO_ENABLED=0 go build
		docker build -t cgi .
run:
		docker run -p 1234:1234 -e MICRO_REGISTRY=mdns cgi