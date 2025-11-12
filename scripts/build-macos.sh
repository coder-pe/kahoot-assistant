#!/bin/bash

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║             Compilando para macOS                              ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Ir al directorio raíz del proyecto
cd "$(dirname "$0")/.." || exit 1

# Crear directorio builds si no existe
mkdir -p builds

# Determinar arquitectura a compilar
ARCH=${1:-"all"}

compile_intel() {
    echo "[1/2] Compilando para macOS Intel (x64)..."
    GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o builds/kahoot-assistant-darwin-amd64 ./cmd/kahoot-assistant
    if [ $? -eq 0 ]; then
        echo "✅ macOS Intel compilado"
        ls -lh builds/kahoot-assistant-darwin-amd64 | awk '{print "  Tamaño: " $5}'
    else
        echo "❌ Error compilando macOS Intel"
        return 1
    fi
}

compile_arm() {
    echo "[2/2] Compilando para macOS Apple Silicon (M1/M2/M3/M4)..."
    GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o builds/kahoot-assistant-darwin-arm64 ./cmd/kahoot-assistant
    if [ $? -eq 0 ]; then
        echo "✅ macOS Apple Silicon compilado"
        ls -lh builds/kahoot-assistant-darwin-arm64 | awk '{print "  Tamaño: " $5}'
    else
        echo "❌ Error compilando macOS Apple Silicon"
        return 1
    fi
}

case $ARCH in
    intel|amd64)
        compile_intel
        ;;
    arm|arm64|m1|m2|m3|m4)
        compile_arm
        ;;
    all|*)
        compile_intel
        echo ""
        compile_arm
        ;;
esac

echo ""
echo "✅ Compilación completada"
echo ""
echo "Binarios en builds/:"
ls -lh builds/kahoot-assistant-darwin-* 2>/dev/null | awk '{print "  " $9 " - " $5}'
echo ""
echo "Uso:"
echo "  $0           # Compila ambas arquitecturas"
echo "  $0 intel     # Solo Intel"
echo "  $0 arm       # Solo Apple Silicon"
