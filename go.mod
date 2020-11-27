module github.com/mianju802/account

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/mianju802/lib/apollo => ../lib/apollo

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/mianju802/lib/apollo v0.0.0-00010101000000-000000000000
	github.com/mianju802/protocol/service/account v0.0.0-20201119134917-e362dc07d67a
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
)
