# Hello-grpc-gateway

Build according to: [gRPC-Gateway: Tutorial](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/)

## Generate and serve swagger.json

Run:
```bash
go intall github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
protoc -I ./proto --swagger_out=logtostderr=true:. ./proto/helloworld/hello_world.proto
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
swagger serve -F=swagger --no-open --port 65534 ./helloworld/hello_world.swagger.json
```

## Another tutorial by 煎鱼
Tutorial: [Grpc+Grpc Gateway实践二 有些复杂的Hello World](https://segmentfault.com/a/1190000013408485)
Modified code on branch: [another_practice](https://github.com/theoriz0/hello-grpc-gateway/blob/another_practice/README.md)

The path and name of certs has been hard coded to simplify the code.
To run the server and client you need to create server.pem and server.key first. For example, run:
```bash
openssl ecparam -genkey -name secp384r1 -out certs/server.key
openssl req -new -x509 -sha256 -key certs/server.key -out certs/server.pem -days 3650 -addext "subjectAltName = DNS:dev.io"
```
And when asked for ```Common Name (eg, your name or your server's hostname) []:```, use ```dev.io``` since it's the cert name used in main.go and client.go. 

## Appendix
About ```-addext "subjectAltName = ..."``` 
[GENERAL: What should I do if I get an “x509: certificate relies on legacy Common Name field” error?](https://jfrog.com/knowledge-base/general-what-should-i-do-if-i-get-an-x509-certificate-relies-on-legacy-common-name-field-error/)
[stackexchange](https://security.stackexchange.com/questions/74345/provide-subjectaltname-to-openssl-directly-on-the-command-line)
