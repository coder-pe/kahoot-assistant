@echo off
echo ============================================
echo   Kahoot Assistant - Script de Compilacion
echo ============================================
echo.

echo Verificando dependencias...
go mod tidy
if errorlevel 1 (
    echo Error: No se pudieron descargar las dependencias
    pause
    exit /b 1
)

echo.
echo Descargando dependencias...
go mod download
if errorlevel 1 (
    echo Error: No se pudieron descargar las dependencias
    pause
    exit /b 1
)

echo.
echo Compilando para Windows 64 bits...
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o kahoot-assistant.exe .
if errorlevel 1 (
    echo Error: Fallo la compilacion
    pause
    exit /b 1
)

echo.
echo ============================================
echo   Compilacion exitosa!
echo   Ejecutable: kahoot-assistant.exe
echo ============================================
echo.
echo Para ejecutar el programa:
echo   .\kahoot-assistant.exe
echo.
pause
