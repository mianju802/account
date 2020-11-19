module github.com/mianju802/account

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/mianju802/protocol/service/account v0.0.0-20201119134917-e362dc07d67a
	github.com/micro/go-micro v1.18.0
)
