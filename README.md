# Kahoot Assistant

Programa multi-plataforma en Go que captura pantalla o portapapeles, extrae texto con OCR (si es imagen), consulta a Gemini AI y envía respuestas por email y/o consola.

## 🚀 Inicio Rápido

**¿Primera vez? Lee esto:**

1. **[docs/QUICKSTART.md](docs/QUICKSTART.md)** - Guía rápida de 5 minutos
2. **[docs/README_MULTIPLATFORM.md](docs/README_MULTIPLATFORM.md)** - Guía completa multi-plataforma

## 📦 Plataformas Soportadas

- ✅ Windows 64 bits
- ✅ macOS Intel y Apple Silicon (M1/M2/M3/M4)
- ✅ Linux 64 bits y ARM64

## 📥 Instalación

### Descargar Binarios

Los binarios compilados están en `builds/` después de compilar:

- `builds/kahoot-assistant-windows-amd64.exe` - Windows 64 bits
- `builds/kahoot-assistant-darwin-arm64` - macOS Apple Silicon
- `builds/kahoot-assistant-darwin-amd64` - macOS Intel
- `builds/kahoot-assistant-linux-amd64` - Linux 64 bits

### Compilar desde Código

```bash
# Compilar para todas las plataformas
./scripts/build-all.sh

# Compilar solo para tu plataforma actual
go build -o kahoot-assistant ./cmd/kahoot-assistant
```

## 🛠️ Configuración

1. Copia el archivo de ejemplo:
```bash
cp config.example.yaml config.yaml
```

2. Edita `config.yaml` con tus datos:
- Token de Gemini API: https://makersuite.google.com/app/apikey
- Credenciales de Gmail (opcional)

3. Instala Tesseract OCR:
- **Windows**: https://github.com/UB-Mannheim/tesseract/wiki
- **macOS**: `brew install tesseract`
- **Linux**: `sudo apt install tesseract-ocr`

## 📚 Documentación Completa

Toda la documentación está en el directorio `docs/`:

- **[docs/INDEX.md](docs/INDEX.md)** - Índice completo de documentación
- **[docs/QUICKSTART.md](docs/QUICKSTART.md)** - Inicio rápido
- **[docs/README_MULTIPLATFORM.md](docs/README_MULTIPLATFORM.md)** - Guía multi-plataforma
- **[docs/INSTALL.md](docs/INSTALL.md)** - Instalación paso a paso
- **[docs/README_ES.md](docs/README_ES.md)** - README en español
- **[docs/PROJECT_SUMMARY.md](docs/PROJECT_SUMMARY.md)** - Resumen técnico

## 🎯 Características

- ✅ **Doble modo de captura:**
  - **Screenshot**: Captura de pantalla (Print Screen en Windows, Cmd+Shift+4 en macOS)
  - **Clipboard**: Lectura de portapapeles con Ctrl+C (Windows/Linux) o Cmd+C (macOS)
- ✅ Soporta texto e imágenes en el portapapeles
- ✅ OCR con Tesseract para imágenes
- ✅ Integración con Gemini AI (modelo configurable)
- ✅ Respuestas en consola
- ✅ Envío opcional por Gmail
- ✅ Sistema de logging
- ✅ Procesamiento < 15 segundos

## 📂 Estructura del Proyecto

```
kahoot-assistant/
├── cmd/
│   └── kahoot-assistant/    # Programa principal
│       └── main.go
├── internal/                # Paquetes internos
│   ├── config/              # Configuración
│   ├── screenshot/          # Captura de pantalla
│   ├── clipboard/           # Lectura de portapapeles (NUEVO)
│   ├── keyboard/            # Detección de teclas (Print Screen / Ctrl+C)
│   ├── ocr/                 # OCR con Tesseract
│   ├── ai/                  # Cliente Gemini
│   ├── notification/        # Envío de emails
│   ├── logger/              # Sistema de logging
│   └── processor/           # Procesamiento principal
├── docs/                    # Documentación
├── scripts/                 # Scripts de compilación
├── builds/                  # Binarios compilados
├── config.yaml              # Tu configuración
├── config.example.yaml      # Plantilla de configuración
└── README.md                # Este archivo
```

## 💻 Uso

### Windows
```cmd
builds\kahoot-assistant-windows-amd64.exe
```
**Dos modos disponibles:**
- **Modo Screenshot**: Presiona **Print Screen** para capturar la pantalla
- **Modo Clipboard**: Copia texto/imagen con **Ctrl+C** y el programa procesará automáticamente

### macOS
```bash
./builds/kahoot-assistant-darwin-arm64
```
**Elige tu modo de captura:**
1. **Screenshot (s)**:
   - Toma screenshot con Cmd+Shift+4 (área) o Cmd+Shift+3 (pantalla)
   - Escribe `s` y presiona Enter
2. **Clipboard (c)**:
   - Copia el texto/imagen con Cmd+C
   - Escribe `c` y presiona Enter

### Linux
```bash
./builds/kahoot-assistant-linux-amd64
```
**Elige tu modo de captura:**
1. **Screenshot (s)**:
   - Toma screenshot con Print Screen
   - Escribe `s` y presiona Enter
2. **Clipboard (c)**:
   - Copia el texto/imagen con Ctrl+C
   - Escribe `c` y presiona Enter

## 🔗 Recursos

- **Gemini API**: https://makersuite.google.com/app/apikey
- **Tesseract**: Ver documentación por plataforma
- **App Password Gmail**: https://myaccount.google.com/apppasswords

## 📖 Más Información

Consulta la **[documentación completa](docs/INDEX.md)** en el directorio `docs/`.

## ⚖️ Licencia

Código abierto para uso educativo. Úsalo responsablemente.

---

**Versión:** 2.0.0 Multi-Plataforma + Clipboard
**Estado:** ✅ Listo para producción
**Novedades v2.0.0:**
- ✨ Soporte para portapapeles (Ctrl+C / Cmd+C)
- ✨ Detección automática de texto e imágenes en clipboard
- ✨ Procesamiento directo de texto sin necesidad de OCR
