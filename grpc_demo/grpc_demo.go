
syntax = "proto3";
// 定义包名
package grpc_demo;

// 定义Greeter服务
service Greeter {
// 定义SayHello方法，接受HelloRequest消息， 并返回HelloReply消息
rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// 定义HelloRequest消息
message HelloRequest {
// name字段
string name = 1;
}

// 定义HelloReply消息
message HelloReply {
// message字段
string message = 1;
}