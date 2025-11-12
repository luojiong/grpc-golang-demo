package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc-demo/gen/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// AddRequest 定义HTTP请求结构
type AddRequest struct {
	A int32 `json:"a" binding:"required"`
	B int32 `json:"b" binding:"required"`
}

// AddResponse 定义HTTP响应结构
type AddResponse struct {
	Result int32 `json:"result"`
}

var grpcClient proto.CalculatorClient

func main() {
	// 连接到gRPC服务器
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpcClient = proto.NewCalculatorClient(conn)

	// 创建 Gin 引擎
	router := gin.Default()

	// CORS 中间件
	router.Use(corsMiddleware())

	// 静态文件服务
	router.Static("/", ".")

	// API 路由
	router.POST("/api/add", handleAdd)

	fmt.Println("HTTP Gateway listening on :8080")
	fmt.Println("访问 http://localhost:8080/test.html 来测试")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// corsMiddleware CORS 中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

// handleAdd 处理加法请求
func handleAdd(c *gin.Context) {
	var req AddRequest

	// 解析并验证请求
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Invalid request: %v", err),
		})
		return
	}

	// 调用gRPC服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	grpcResp, err := grpcClient.Add(ctx, &proto.AddRequest{A: req.A, B: req.B})
	if err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("gRPC error: %v", err),
		})
		return
	}

	// 返回结果
	c.JSON(200, AddResponse{Result: grpcResp.Result})
}
