@echo off
echo ========================================
echo   测试 HTTP Gateway API
echo ========================================
echo.
echo 测试加法: 100 + 200
echo.

curl -X POST http://localhost:8080/api/add ^
  -H "Content-Type: application/json" ^
  -d "{\"a\": 100, \"b\": 200}"

echo.
echo.
echo 预期结果: {"result":300}
echo.

pause

