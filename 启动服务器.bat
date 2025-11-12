@echo off
echo ========================================
echo   启动 gRPC 服务器 (端口 50051)
echo ========================================
echo.
echo 按 Ctrl+C 停止服务器
echo ========================================
echo.

go run server/main.go

pause

