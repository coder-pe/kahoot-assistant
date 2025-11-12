#!/bin/bash

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║             Compilando para Linux                              ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Ir al directorio raíz del proyecto
cd "$(dirname "$0")/.." || exit 1

# Crear directorio builds si no existe
mkdir -p builds

# Determinar arquitectura a compilar
ARCH=${1:-"all"}

compile_amd64() {
    echo "[1/2] Compilando para Linux 64 bits (amd64)..."
    GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o builds/kahoot-assistant-linux-amd64 ./cmd/kahoot-assistant
    if [ $? -eq 0 ]; then
        echo "✅ Linux amd64 compilado"
        ls -lh builds/kahoot-assistant-linux-amd64 | awk '{print "  Tamaño: " $5}'
    else
        echo "❌ Error compilando Linux amd64"
        return 1
    fi
}

compile_arm64() {
    echo "[2/2] Compilando para Linux ARM64 (Raspberry Pi, etc.)..."
    GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o builds/kahoot-assistant-linux-arm64 ./cmd/kahoot-assistant
    if [ $? -eq 0 ]; then
        echo "✅ Linux ARM64 compilado"
        ls -lh builds/kahoot-assistant-linux-arm64 | awk '{print "  Tamaño: " $5}'
    else
        echo "❌ Error compilando Linux ARM64"
        return 1
    fi
}

case $ARCH in
    amd64|x64)
        compile_amd64
        ;;
    arm64|arm|rpi)
        compile_arm64
        ;;
    all|*)
        compile_amd64
        echo ""
        compile_arm64
        ;;
esac

echo ""
echo "✅ Compilación completada"
echo ""
echo "Binarios en builds/:"
ls -lh builds/kahoot-assistant-linux-* 2>/dev/null | awk '{print "  " $9 " - " $5}'
echo ""
echo "Uso:"
echo "  $0           # Compila ambas arquitecturas"
echo "  $0 amd64     # Solo x64"
echo "  $0 arm64     # Solo ARM64"
