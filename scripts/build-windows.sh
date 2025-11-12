#!/bin/bash

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║          Compilando para Windows 64 bits                       ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Ir al directorio raíz del proyecto
cd "$(dirname "$0")/.." || exit 1

# Crear directorio builds si no existe
mkdir -p builds

echo "Compilando kahoot-assistant para Windows..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o builds/kahoot-assistant-windows-amd64.exe ./cmd/kahoot-assistant

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Compilación exitosa!"
    echo ""
    ls -lh builds/kahoot-assistant-windows-amd64.exe | awk '{print "Archivo: " $9 "\nTamaño: " $5}'
    echo ""
    echo "Ubicación: builds/kahoot-assistant-windows-amd64.exe"
else
    echo ""
    echo "❌ Error en la compilación"
    exit 1
fi
