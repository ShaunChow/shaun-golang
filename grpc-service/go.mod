module github.com/shaun-golang/grpc-service

go 1.14

replace github.com/shaun-golang/micro-service => ../micro-service

require (
	github.com/shaun-golang/micro-service v1.0.0
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	google.golang.org/grpc v1.34.0
)
