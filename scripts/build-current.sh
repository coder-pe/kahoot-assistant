#!/bin/bash

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║        Compilando para la plataforma actual                    ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Ir al directorio raíz del proyecto
cd "$(dirname "$0")/.." || exit 1

# Crear directorio builds si no existe
mkdir -p builds

# Detectar plataforma y arquitectura actual
OS=$(uname -s)
ARCH=$(uname -m)

# Convertir a nombres de Go
case $OS in
    Darwin)
        GOOS="darwin"
        OS_NAME="macOS"
        ;;
    Linux)
        GOOS="linux"
        OS_NAME="Linux"
        ;;
    MINGW*|MSYS*|CYGWIN*)
        GOOS="windows"
        OS_NAME="Windows"
        ;;
    *)
        echo "❌ Sistema operativo no soportado: $OS"
        exit 1
        ;;
esac

case $ARCH in
    x86_64|amd64)
        GOARCH="amd64"
        ARCH_NAME="x64"
        ;;
    arm64|aarch64)
        GOARCH="arm64"
        ARCH_NAME="ARM64"
        ;;
    *)
        echo "❌ Arquitectura no soportada: $ARCH"
        exit 1
        ;;
esac

echo "Plataforma detectada: $OS_NAME $ARCH_NAME"
echo "Compilando para: GOOS=$GOOS GOARCH=$GOARCH"
echo ""

# Determinar nombre del binario
if [ "$GOOS" = "windows" ]; then
    BINARY="builds/kahoot-assistant-${GOOS}-${GOARCH}.exe"
else
    BINARY="builds/kahoot-assistant-${GOOS}-${GOARCH}"
fi

echo "Compilando..."
GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$BINARY" ./cmd/kahoot-assistant

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Compilación exitosa!"
    echo ""
    ls -lh "$BINARY" | awk '{print "Archivo: " $9 "\nTamaño: " $5}'
    echo ""
    echo "Para ejecutar:"
    if [ "$GOOS" = "windows" ]; then
        echo "  $BINARY"
    else
        echo "  ./$BINARY"
    fi
else
    echo ""
    echo "❌ Error en la compilación"
    exit 1
fi
