# Kahoot Assistant

Programa multiplataforma en Go que captura pantalla, extrae texto con OCR, consulta a Gemini AI y envía respuestas por email y/o consola.

## 🌍 Plataformas Soportadas

- ✅ **Windows** 64 bits (detección automática de Print Screen)
- ✅ **macOS** Intel y Apple Silicon (M1/M2/M3/M4)
- ✅ **Linux** 64 bits y ARM64

## 📥 Descarga Rápida

Descarga el binario para tu sistema:

| Sistema Operativo | Binario | Tamaño |
|-------------------|---------|--------|
| Windows 64 bits | `kahoot-assistant-windows-amd64.exe` | 15 MB |
| macOS Intel | `kahoot-assistant-darwin-amd64` | 14 MB |
| macOS M1/M2/M3/M4 | `kahoot-assistant-darwin-arm64` | 14 MB |
| Linux 64 bits | `kahoot-assistant-linux-amd64` | 14 MB |
| Linux ARM64 | `kahoot-assistant-linux-arm64` | 13 MB |

## 🚀 Inicio Rápido

### Windows
```cmd
1. Descarga kahoot-assistant-windows-amd64.exe
2. Instala Tesseract OCR
3. Configura config.yaml con tu token de Gemini
4. Ejecuta el programa
5. Presiona Print Screen cuando veas una pregunta de Kahoot
```

### macOS
```bash
1. Descarga kahoot-assistant-darwin-arm64 (o amd64 para Intel)
2. chmod +x kahoot-assistant-darwin-arm64
3. brew install tesseract
4. Configura config.yaml con tu token de Gemini
5. Ejecuta el programa
6. Toma screenshot (Cmd+Shift+3/4) y presiona Enter
```

### Linux
```bash
1. Descarga kahoot-assistant-linux-amd64
2. chmod +x kahoot-assistant-linux-amd64
3. sudo apt install tesseract-ocr  # o tu gestor de paquetes
4. Configura config.yaml con tu token de Gemini
5. Ejecuta el programa
6. Toma screenshot y presiona Enter
```

## 🎯 Características

- ✅ Captura de pantalla (automática en Windows, manual en macOS/Linux)
- ✅ OCR con Tesseract (extracción de texto)
- ✅ Consulta a Gemini AI
- ✅ Respuesta en consola
- ✅ Envío opcional por Gmail
- ✅ Sistema de logging
- ✅ Procesamiento < 15 segundos
- ✅ Configuración flexible

## 📚 Documentación

- **Guía Multi-Plataforma:** `README_MULTIPLATFORM.md` ⭐
- **Inicio Rápido:** `QUICKSTART.md`
- **Instalación Detallada:** `INSTALL.md`
- **Compilar desde macOS:** `README_macOS.md`
- **Resumen Técnico:** `PROJECT_SUMMARY.md`

## ⚙️ Configuración

### 1. Crear archivo de configuración

```bash
cp config.example.yaml config.yaml
```

### 2. Obtener Token de Gemini

https://makersuite.google.com/app/apikey

### 3. Editar config.yaml

```yaml
gemini_api_key: "TU_TOKEN_AQUI"

email:
  enabled: true  # o false si no quieres email
  from: "tu_email@gmail.com"
  password: "tu_app_password"
  to: "destino@gmail.com"

output:
  console: true
  email: true
  log_file: true
```

## 🛠️ Requisitos

### Todas las plataformas
- **Tesseract OCR** (para extracción de texto)
- **Token de Gemini API** (para respuestas IA)
- **Conexión a Internet**

### Instalación de Tesseract

**Windows:**
```
https://github.com/UB-Mannheim/tesseract/wiki
```

**macOS:**
```bash
brew install tesseract
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt install tesseract-ocr
```

**Linux (Fedora):**
```bash
sudo dnf install tesseract
```

**Linux (Arch):**
```bash
sudo pacman -S tesseract
```

## 💻 Uso

### En Windows

```cmd
.\kahoot-assistant-windows-amd64.exe
```

Presiona **Print Screen** cuando aparezca una pregunta.

### En macOS

```bash
./kahoot-assistant-darwin-arm64
```

1. Toma screenshot: **Cmd+Shift+3** (pantalla completa) o **Cmd+Shift+4** (selección)
2. Presiona **Enter** en el programa

### En Linux

```bash
./kahoot-assistant-linux-amd64
```

1. Toma screenshot: **Print Screen** o usa `gnome-screenshot`, `flameshot`, etc.
2. Presiona **Enter** en el programa

## 🔨 Compilar desde Código Fuente

### Compilar para TODAS las plataformas

```bash
./build-all.sh
```

Genera binarios para:
- Windows 64 bits
- macOS Intel y Apple Silicon
- Linux 64 bits y ARM64

### Compilar solo para tu plataforma

```bash
go build -o kahoot-assistant
```

### Compilación cruzada

```bash
# Para Windows desde macOS/Linux
GOOS=windows GOARCH=amd64 go build -o kahoot-assistant.exe

# Para macOS M1/M2/M3/M4 desde Windows/Linux
GOOS=darwin GOARCH=arm64 go build -o kahoot-assistant-mac

# Para Linux desde Windows/macOS
GOOS=linux GOARCH=amd64 go build -o kahoot-assistant-linux
```

## 📊 Arquitectura

```
Usuario → Captura/Screenshot → OCR → Gemini AI → Salida
                                                    ├─ Consola
                                                    ├─ Email
                                                    └─ Log
```

### Archivos por Plataforma

- **Común:** `main.go`, `config.go`, `gemini.go`, `email.go`, `logger.go`, `ocr.go`
- **Windows:** `keyboard_windows.go`, `screenshot_windows.go`, `process_windows.go`
- **macOS:** `keyboard_darwin.go`, `screenshot_darwin.go`, `process_darwin.go`
- **Linux:** `keyboard_linux.go`, `screenshot_linux.go`, `process_linux.go`

## 🔒 Seguridad

⚠️ **IMPORTANTE:**
- Nunca compartas `config.yaml` (contiene tokens y contraseñas)
- Usa App Password de Gmail, NO tu contraseña normal
- Los tokens de Gemini son personales

## 🐛 Solución de Problemas

### "Tesseract not found"
```bash
# Verifica instalación
tesseract --version

# Reinstala si es necesario
```

### "Invalid API key"
- Verifica que copiaste correctamente el token
- Genera un nuevo token si es necesario

### macOS: "No se puede abrir el programa"
```bash
xattr -d com.apple.quarantine kahoot-assistant-darwin-arm64
```

### Linux: "Permission denied"
```bash
chmod +x kahoot-assistant-linux-amd64
```

## 📈 Rendimiento

- **Tiempo total:** 5-15 segundos
- **OCR:** 2-4 segundos
- **Gemini:** 3-5 segundos
- **Email:** 1-2 segundos

## 🌟 Diferencias por Plataforma

| Característica | Windows | macOS | Linux |
|----------------|---------|-------|-------|
| Detección automática de tecla | ✅ Print Screen | ❌ Manual | ❌ Manual |
| Captura automática | ✅ | ❌ | ❌ |
| Método | Automático | Cmd+Shift+3/4 → Enter | Print Screen → Enter |

## 📦 Dependencias de Go

```go
github.com/google/generative-ai-go  // Gemini API
github.com/kbinani/screenshot       // Captura (Windows)
golang.org/x/sys                    // APIs del sistema
google.golang.org/api               // Google APIs
gopkg.in/yaml.v3                    // Configuración
```

## 🔗 Recursos

- **Gemini API:** https://makersuite.google.com/app/apikey
- **Tesseract:**
  - Windows: https://github.com/UB-Mannheim/tesseract/wiki
  - macOS: `brew install tesseract`
  - Linux: Gestor de paquetes de tu distro
- **App Password Gmail:** https://myaccount.google.com/apppasswords

## 📖 Guías Específicas

- **¿Primera vez?** → Lee `QUICKSTART.md`
- **Instalación paso a paso** → Lee `INSTALL.md`
- **Multi-plataforma** → Lee `README_MULTIPLATFORM.md`
- **Compilar desde macOS** → Lee `README_macOS.md`

## ⚖️ Licencia

Código abierto para uso educativo.

**Úsalo responsablemente:**
- No para hacer trampa en exámenes reales
- Solo para práctica y aprendizaje
- Respeta las políticas de tu institución

## 🎉 Estado

✅ **Listo para producción**
- Windows ✅
- macOS ✅
- Linux ✅

Versión: 1.0.0 Multi-Plataforma
