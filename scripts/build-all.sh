#!/bin/bash

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║     Kahoot Assistant - Compilación Multi-Plataforma           ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Colores para output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Ir al directorio raíz del proyecto
cd "$(dirname "$0")/.." || exit 1

# Limpiar builds anteriores
echo -e "${BLUE}Limpiando builds anteriores...${NC}"
rm -rf builds/*
mkdir -p builds

# Verificar dependencias
echo -e "${BLUE}Verificando dependencias...${NC}"
go mod tidy
if [ $? -ne 0 ]; then
    echo "Error: No se pudieron verificar las dependencias"
    exit 1
fi

echo ""
echo -e "${BLUE}Descargando dependencias...${NC}"
go mod download
if [ $? -ne 0 ]; then
    echo "Error: No se pudieron descargar las dependencias"
    exit 1
fi

echo ""
echo "════════════════════════════════════════════════════════════════"
echo "Compilando para múltiples plataformas..."
echo "════════════════════════════════════════════════════════════════"
echo ""

# Compilar para Windows 64 bits
echo -e "${YELLOW}[1/5] Compilando para Windows 64 bits...${NC}"
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o builds/kahoot-assistant-windows-amd64.exe ./cmd/kahoot-assistant
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Windows 64 bits compilado${NC}"
    ls -lh builds/kahoot-assistant-windows-amd64.exe | awk '{print "  Tamaño: " $5}'
else
    echo "✗ Error compilando Windows 64 bits"
fi
echo ""

# Compilar para macOS Intel
echo -e "${YELLOW}[2/5] Compilando para macOS Intel (x64)...${NC}"
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o builds/kahoot-assistant-darwin-amd64 ./cmd/kahoot-assistant
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ macOS Intel compilado${NC}"
    ls -lh builds/kahoot-assistant-darwin-amd64 | awk '{print "  Tamaño: " $5}'
else
    echo "✗ Error compilando macOS Intel"
fi
echo ""

# Compilar para macOS Apple Silicon
echo -e "${YELLOW}[3/5] Compilando para macOS Apple Silicon (M1/M2/M3/M4)...${NC}"
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o builds/kahoot-assistant-darwin-arm64 ./cmd/kahoot-assistant
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ macOS Apple Silicon compilado${NC}"
    ls -lh builds/kahoot-assistant-darwin-arm64 | awk '{print "  Tamaño: " $5}'
else
    echo "✗ Error compilando macOS Apple Silicon"
fi
echo ""

# Compilar para Linux 64 bits
echo -e "${YELLOW}[4/5] Compilando para Linux 64 bits...${NC}"
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o builds/kahoot-assistant-linux-amd64 ./cmd/kahoot-assistant
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Linux 64 bits compilado${NC}"
    ls -lh builds/kahoot-assistant-linux-amd64 | awk '{print "  Tamaño: " $5}'
else
    echo "✗ Error compilando Linux 64 bits"
fi
echo ""

# Compilar para Linux ARM64 (Raspberry Pi, etc.)
echo -e "${YELLOW}[5/5] Compilando para Linux ARM64...${NC}"
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o builds/kahoot-assistant-linux-arm64 ./cmd/kahoot-assistant
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Linux ARM64 compilado${NC}"
    ls -lh builds/kahoot-assistant-linux-arm64 | awk '{print "  Tamaño: " $5}'
else
    echo "✗ Error compilando Linux ARM64"
fi
echo ""

# Resumen
echo "════════════════════════════════════════════════════════════════"
echo -e "${GREEN}✓ Compilación completada${NC}"
echo "════════════════════════════════════════════════════════════════"
echo ""
echo "Binarios generados en builds/:"
echo ""
ls -lh builds/kahoot-assistant-* 2>/dev/null | awk '{print "  " $9 " - " $5}'
echo ""
