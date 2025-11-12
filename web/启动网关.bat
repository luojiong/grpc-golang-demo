@echo off
echo ========================================
echo   启动 HTTP Gateway (端口 8080)
echo ========================================
echo.
echo 请确保 gRPC 服务器已经在运行 (端口 50051)
echo.
echo 启动后请访问: http://localhost:8080/test.html
echo.
echo 按 Ctrl+C 停止服务器
echo ========================================
echo.

cd ..
go run web/gateway.go

pause

