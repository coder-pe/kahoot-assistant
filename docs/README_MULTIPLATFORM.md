# Kahoot Assistant - Guía Multi-Plataforma

Este programa ahora funciona en **Windows**, **macOS** y **Linux**!

## Diferencias por Plataforma

### Windows
- ✅ Detección automática de tecla **Print Screen**
- ✅ Captura automática de pantalla
- 🎯 Experiencia más fluida

### macOS
- ℹ️  Toma screenshot manualmente (Cmd+Shift+3/4/5)
- ℹ️  Presiona **Enter** en el programa para procesar
- ℹ️  Busca el screenshot más reciente en el Escritorio

### Linux
- ℹ️  Toma screenshot manualmente (Print Screen, gnome-screenshot, etc.)
- ℹ️  Presiona **Enter** en el programa para procesar
- ℹ️  Busca screenshots en ~/Pictures/Screenshots y otras ubicaciones

## Binarios Disponibles

```
kahoot-assistant-windows-amd64.exe  → Windows 64 bits
kahoot-assistant-darwin-amd64       → macOS Intel (x64)
kahoot-assistant-darwin-arm64       → macOS Apple Silicon (M1/M2/M3/M4)
kahoot-assistant-linux-amd64        → Linux 64 bits
kahoot-assistant-linux-arm64        → Linux ARM64 (Raspberry Pi, etc.)
```

## Instalación por Plataforma

### 🪟 Windows

#### 1. Descargar
- `kahoot-assistant-windows-amd64.exe`
- `config.example.yaml`

#### 2. Instalar Tesseract OCR
```
https://github.com/UB-Mannheim/tesseract/wiki
```

#### 3. Configurar
Renombra `config.example.yaml` a `config.yaml` y completa tus datos.

#### 4. Ejecutar
```cmd
.\kahoot-assistant-windows-amd64.exe
```

#### 5. Usar
1. Abre Kahoot en navegador
2. Presiona **Print Screen** cuando aparezca una pregunta
3. Ve la respuesta en 5-15 segundos

---

### 🍎 macOS

#### 1. Descargar
- `kahoot-assistant-darwin-arm64` (para M1/M2/M3/M4)
- `kahoot-assistant-darwin-amd64` (para Mac Intel)
- `config.example.yaml`

#### 2. Dar permisos de ejecución
```bash
chmod +x kahoot-assistant-darwin-arm64
```

#### 3. Instalar Tesseract OCR
```bash
brew install tesseract
```

Si no tienes Homebrew:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

#### 4. Configurar
```bash
cp config.example.yaml config.yaml
nano config.yaml  # o usa TextEdit
```

Completa tus tokens y credenciales.

#### 5. Ejecutar
```bash
./kahoot-assistant-darwin-arm64  # o darwin-amd64 para Intel
```

#### 6. Usar
1. Ejecuta el programa
2. Abre Kahoot en navegador
3. Cuando aparezca una pregunta, toma screenshot:
   - **Cmd+Shift+3** = Pantalla completa
   - **Cmd+Shift+4** = Selección de área
   - **Cmd+Shift+5** = Herramienta de captura
4. Presiona **Enter** en el programa
5. Ve la respuesta en 5-15 segundos

**Nota:** El screenshot se guarda automáticamente en tu Escritorio.

---

### 🐧 Linux

#### 1. Descargar
- `kahoot-assistant-linux-amd64`
- `config.example.yaml`

#### 2. Dar permisos de ejecución
```bash
chmod +x kahoot-assistant-linux-amd64
```

#### 3. Instalar Tesseract OCR

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install tesseract-ocr
```

**Fedora:**
```bash
sudo dnf install tesseract
```

**Arch:**
```bash
sudo pacman -S tesseract
```

#### 4. Configurar
```bash
cp config.example.yaml config.yaml
nano config.yaml  # o tu editor preferido
```

Completa tus tokens y credenciales.

#### 5. Ejecutar
```bash
./kahoot-assistant-linux-amd64
```

#### 6. Usar
1. Ejecuta el programa
2. Abre Kahoot en navegador
3. Cuando aparezca una pregunta, toma screenshot:
   - **Print Screen** = Pantalla completa
   - **Alt+Print Screen** = Ventana activa
   - **Shift+Print Screen** = Selección
   - O usa: `gnome-screenshot`, `flameshot`, `spectacle`, etc.
4. Presiona **Enter** en el programa
5. Ve la respuesta en 5-15 segundos

**Ubicaciones de screenshots buscadas:**
- `~/Pictures/Screenshots/`
- `~/Pictures/`
- `~/Desktop/`

---

## Configuración Común (Todas las Plataformas)

### 1. Obtener Token de Gemini API
```
https://makersuite.google.com/app/apikey
```

### 2. Obtener App Password de Gmail (opcional)
```
https://myaccount.google.com/apppasswords
```

### 3. Editar config.yaml

```yaml
gemini_api_key: "TU_TOKEN_AQUI"

email:
  enabled: true  # o false
  from: "tu_email@gmail.com"
  password: "tu_app_password_16_caracteres"
  to: "destino@gmail.com"

output:
  console: true
  email: true
  log_file: true

ocr:
  language: "eng"
  tesseract_path: ""  # Dejar vacío para PATH del sistema

timeouts:
  max_processing_time: 15
  gemini_timeout: 10
  email_timeout: 5

log_file_path: "kahoot_log.txt"
```

## Compilar desde Código Fuente

### Compilar para TODAS las plataformas
```bash
./build-all.sh
```

Genera:
- Windows 64 bits
- macOS Intel
- macOS Apple Silicon
- Linux 64 bits
- Linux ARM64

### Compilar solo para tu plataforma

**En macOS/Linux:**
```bash
go build -o kahoot-assistant
```

**En Windows:**
```cmd
go build -o kahoot-assistant.exe
```

### Compilación cruzada manual

**Desde cualquier OS a Windows:**
```bash
GOOS=windows GOARCH=amd64 go build -o kahoot-assistant.exe
```

**Desde cualquier OS a macOS (M1/M2/M3/M4):**
```bash
GOOS=darwin GOARCH=arm64 go build -o kahoot-assistant-mac
```

**Desde cualquier OS a Linux:**
```bash
GOOS=linux GOARCH=amd64 go build -o kahoot-assistant-linux
```

## Comparación de Experiencia de Usuario

| Característica | Windows | macOS | Linux |
|----------------|---------|-------|-------|
| Detección automática Print Screen | ✅ Sí | ❌ No | ❌ No |
| Captura automática | ✅ Sí | ❌ No | ❌ No |
| Método de captura | Automático | Manual | Manual |
| Trigger del programa | Print Screen | Enter | Enter |
| OCR | ✅ | ✅ | ✅ |
| Gemini API | ✅ | ✅ | ✅ |
| Email | ✅ | ✅ | ✅ |
| Logging | ✅ | ✅ | ✅ |

## Solución de Problemas

### macOS: "No se puede abrir porque proviene de un desarrollador no identificado"

```bash
xattr -d com.apple.quarantine kahoot-assistant-darwin-arm64
```

O:
1. Click derecho > Abrir
2. Click en "Abrir" en el diálogo de seguridad

### Linux: "Permission denied"

```bash
chmod +x kahoot-assistant-linux-amd64
```

### Todas las plataformas: "Tesseract not found"

Verifica la instalación:
```bash
tesseract --version
```

Si no aparece, reinstala Tesseract para tu plataforma.

### macOS/Linux: "No se encontraron screenshots"

- Verifica que tomaste el screenshot
- El screenshot debe tener menos de 5 minutos
- En macOS: debe estar en el Escritorio
- En Linux: debe estar en ~/Pictures/Screenshots o ~/Pictures

## Rendimiento por Plataforma

Todos los sistemas tienen rendimiento similar:

- **Tiempo total:** 5-15 segundos
- **Captura:** < 1 segundo (Windows) o instantánea (macOS/Linux)
- **OCR:** 2-4 segundos
- **Gemini:** 3-5 segundos
- **Email:** 1-2 segundos

## Arquitectura Multi-Plataforma

El programa usa **build tags** de Go para compilar código específico por plataforma:

```
main.go                  → Código común
config.go                → Código común
gemini.go                → Código común
email.go                 → Código común
logger.go                → Código común
ocr.go                   → Código común
process_common.go        → Código común

keyboard_windows.go      → Solo Windows
keyboard_darwin.go       → Solo macOS
keyboard_linux.go        → Solo Linux

screenshot_windows.go    → Solo Windows
screenshot_darwin.go     → Solo macOS
screenshot_linux.go      → Solo Linux

process_windows.go       → Solo Windows
process_darwin.go        → Solo macOS
process_linux.go         → Solo Linux
```

## Recursos por Plataforma

### Windows
- Token Gemini: https://makersuite.google.com/app/apikey
- Tesseract: https://github.com/UB-Mannheim/tesseract/wiki
- App Password: https://myaccount.google.com/apppasswords

### macOS
- Token Gemini: https://makersuite.google.com/app/apikey
- Homebrew: https://brew.sh/
- Tesseract: `brew install tesseract`
- App Password: https://myaccount.google.com/apppasswords

### Linux
- Token Gemini: https://makersuite.google.com/app/apikey
- Tesseract: Gestor de paquetes de tu distro
- App Password: https://myaccount.google.com/apppasswords

## Siguiente Paso

Consulta la documentación específica de tu plataforma:
- **Windows:** `README.md` o `INSTALL.md`
- **macOS:** `README_macOS.md`
- **Linux:** Esta guía y `README.md`
- **Inicio rápido:** `QUICKSTART.md`

---

**¡El programa ahora funciona en Windows, macOS y Linux!** 🎉
