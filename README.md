# grpc example
- grpc golang中使用, 包含grpc-gateway
- 自动生成pb.go文件：安装vscode插件 vscode-proto3
- 手动生成文件：protoc --go_out=plugins=grpc:proto proto/helloworld.proto

# run
- go run main.go server
- go run main.go client
- using http: [http://localhost:8081/hello/world](http://localhost:8081/hello/world)

# build server
docker build -t grpcweb/server -f Dockerfile .
docker run --name grpc-server -d -p 9999:9999 grpcweb/server

# build proxy envoy
docker build -t grpcweb/envoy -f Dockerfile.envoy .
docker run --name grpc-envoy -d -p 8080:8080 -p 9901:9901 --link grpc-server:grpc-server grpcweb/envoy
