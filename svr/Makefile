build:
		GOOS=linux CGO_ENABLED=0 go build
		docker build -t svr .
		docker tag svr xufongnian/svr:v1
		docker push xufongnian/svr:v1
run:
		docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns svr