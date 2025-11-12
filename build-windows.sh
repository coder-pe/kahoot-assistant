#!/bin/bash

echo "============================================"
echo "  Kahoot Assistant - Compilación Cruzada"
echo "  macOS -> Windows 64 bits"
echo "============================================"
echo ""

echo "Verificando dependencias..."
go mod tidy
if [ $? -ne 0 ]; then
    echo "Error: No se pudieron verificar las dependencias"
    exit 1
fi

echo ""
echo "Descargando dependencias..."
go mod download
if [ $? -ne 0 ]; then
    echo "Error: No se pudieron descargar las dependencias"
    exit 1
fi

echo ""
echo "Compilando para Windows 64 bits..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o kahoot-assistant.exe .
if [ $? -ne 0 ]; then
    echo "Error: Falló la compilación"
    exit 1
fi

echo ""
echo "============================================"
echo "  ✓ Compilación exitosa!"
echo "  Ejecutable: kahoot-assistant.exe"
echo "============================================"
echo ""
echo "El archivo kahoot-assistant.exe está listo para"
echo "transferir a un sistema Windows 64 bits."
echo ""
echo "Tamaño del ejecutable:"
ls -lh kahoot-assistant.exe | awk '{print "  " $5 " - " $9}'
echo ""
