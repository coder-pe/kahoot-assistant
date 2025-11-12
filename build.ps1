# Script de compilación para Kahoot Assistant (PowerShell)

Write-Host "============================================" -ForegroundColor Cyan
Write-Host "  Kahoot Assistant - Script de Compilacion" -ForegroundColor Cyan
Write-Host "============================================" -ForegroundColor Cyan
Write-Host ""

Write-Host "Verificando dependencias..." -ForegroundColor Yellow
go mod tidy
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: No se pudieron verificar las dependencias" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "Descargando dependencias..." -ForegroundColor Yellow
go mod download
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: No se pudieron descargar las dependencias" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "Compilando para Windows 64 bits..." -ForegroundColor Yellow
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -ldflags="-s -w" -o kahoot-assistant.exe .
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Fallo la compilacion" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "============================================" -ForegroundColor Green
Write-Host "  Compilacion exitosa!" -ForegroundColor Green
Write-Host "  Ejecutable: kahoot-assistant.exe" -ForegroundColor Green
Write-Host "============================================" -ForegroundColor Green
Write-Host ""
Write-Host "Para ejecutar el programa:" -ForegroundColor Cyan
Write-Host "  .\kahoot-assistant.exe" -ForegroundColor White
Write-Host ""
