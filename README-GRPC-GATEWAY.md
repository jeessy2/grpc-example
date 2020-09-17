go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go


- 拷贝`third_party/googleapis/google`到proto目录后, 需要指定引用目录-I./proto
```
protoc -I./proto --go_out=plugins=grpc,paths=source_relative:./service \
proto/helloworld.proto
```

- 生成gw
```
protoc -I./proto --grpc-gateway_out=logtostderr=true,paths=source_relative:./service \
proto/helloworld.proto
```