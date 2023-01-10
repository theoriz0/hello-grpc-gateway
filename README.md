# Hello-grpc-gateway

Build according to: [gRPC-Gateway: Tutorial](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/)

Swagger support added by running:
```bash
go intall github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
protoc -I ./proto --swagger_out=logtostderr=true:. ./proto/helloworld/hello_world.proto
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
swagger serve -F=swagger --no-open --port 65534 ./helloworld/hello_world.swagger.json
```