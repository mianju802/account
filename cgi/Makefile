build:
		GOOS=linux CGO_ENABLED=0 go build
		docker build -t cgi .
		docker tag cgi xufongnian/cgi:v1
		docker push xufongnian/cgi:v1
run:
		docker run -p 1234:1234 -e MICRO_REGISTRY=mdns cgi