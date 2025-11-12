// Package main 是 gRPC Calculator 服务的服务端实现
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-demo/gen/proto" // 导入生成的 protobuf 代码

	"google.golang.org/grpc" // 导入 gRPC 框架
)

// calculatorServer 实现了 Calculator 服务
// 嵌入 UnimplementedCalculatorServer 以满足接口要求
type calculatorServer struct {
	proto.UnimplementedCalculatorServer
}

// Add 方法实现了 Calculator 服务中的 Add RPC 方法
// 功能：对两个整数进行加法运算
// 参数：
//   - ctx: 上下文信息（用于控制请求的生命周期）
//   - req: 请求消息，包含两个被加数 a 和 b
//
// 返回：
//   - 响应消息，包含加法结果
//   - error: 处理过程中的错误信息
func (s *calculatorServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	// 执行加法计算
	result := req.A + req.B
	// 返回加法结果
	return &proto.AddResponse{Result: result}, nil
}

// main 函数是程序入口点，负责启动 gRPC 服务器
func main() {
	// 在 TCP 端口 50051 上监听客户端连接
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建新的 gRPC 服务器实例
	grpcServer := grpc.NewServer()

	// 注册 Calculator 服务到 gRPC 服务器
	proto.RegisterCalculatorServer(grpcServer, &calculatorServer{})

	// 打印服务启动信息
	fmt.Println("Server listening on :50051")

	// 启动 gRPC 服务器，开始接收和处理客户端请求
	// 这是一个阻塞操作，服务器会持续运行直到发生错误
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
