# Compilar para Windows desde macOS

Si estás en una Mac (como MacBook M4) y necesitas compilar el programa para Windows, sigue estos pasos:

## Requisitos en macOS

1. **Go 1.21 o superior**
   ```bash
   # Verificar instalación
   go version

   # Si no está instalado, descarga desde:
   # https://golang.org/dl/
   ```

2. **Git** (opcional, para clonar el repositorio)

## Pasos para Compilar

### 1. Abrir Terminal

Abre Terminal.app en tu Mac

### 2. Navegar al directorio del proyecto

```bash
cd /ruta/a/kahoot-assistant
```

### 3. Ejecutar el script de compilación

```bash
# Dar permisos de ejecución al script (solo la primera vez)
chmod +x build-windows.sh

# Ejecutar la compilación
./build-windows.sh
```

Si todo va bien, verás:
```
============================================
  ✓ Compilación exitosa!
  Ejecutable: kahoot-assistant.exe
============================================

El archivo kahoot-assistant.exe está listo para
transferir a un sistema Windows 64 bits.

Tamaño del ejecutable:
  15M - kahoot-assistant.exe
```

### 4. Transferir el ejecutable a Windows

Ahora tienes `kahoot-assistant.exe` listo para usar en Windows. Puedes transferirlo de varias formas:

**Opción 1: USB**
- Copia el archivo a una memoria USB
- Conecta la USB a tu PC Windows
- Copia el archivo

**Opción 2: Email**
- Envíate el archivo por email
- Descárgalo en Windows
- **Nota:** Algunos servicios de email pueden bloquear archivos .exe

**Opción 3: Cloud Storage (Recomendado)**
- Sube a Google Drive, Dropbox, iCloud, etc.
- Descarga en Windows

**Opción 4: Red local (más rápido)**
```bash
# En macOS, inicia un servidor HTTP simple
python3 -m http.server 8000

# En Windows, abre el navegador y ve a:
# http://[IP-de-tu-Mac]:8000
# Descarga kahoot-assistant.exe
```

## Archivos a Transferir a Windows

Para que el programa funcione en Windows, necesitas transferir:

1. ✅ `kahoot-assistant.exe` - El ejecutable
2. ✅ `config.yaml` - El archivo de configuración (o `config.example.yaml` para crear uno nuevo)
3. ⚠️ **NO** transferir archivos de código fuente (.go) - no son necesarios

## Configuración en Windows

Una vez transferido el archivo a Windows:

### 1. Instalar Tesseract OCR en Windows

Descarga desde: https://github.com/UB-Mannheim/tesseract/wiki
- Ejecuta `tesseract-ocr-w64-setup-5.3.x.exe`
- Marca "Additional language data" > English
- Agrega al PATH: `C:\Program Files\Tesseract-OCR`

### 2. Configurar config.yaml

Crea o edita `config.yaml` en el mismo directorio que el .exe:

```yaml
gemini_api_key: "TU_TOKEN_DE_GEMINI"

email:
  enabled: true
  smtp_host: "smtp.gmail.com"
  smtp_port: 587
  from: "tu_email@gmail.com"
  password: "tu_app_password"
  to: "destino@gmail.com"

output:
  console: true
  email: true
  log_file: true

ocr:
  language: "eng"
  tesseract_path: ""  # Dejar vacío si Tesseract está en PATH

timeouts:
  max_processing_time: 15
  gemini_timeout: 10
  email_timeout: 5

log_file_path: "kahoot_log.txt"
```

### 3. Ejecutar en Windows

**Opción 1: Doble clic**
- Haz doble clic en `kahoot-assistant.exe`

**Opción 2: CMD**
```cmd
kahoot-assistant.exe
```

**Opción 3: PowerShell**
```powershell
.\kahoot-assistant.exe
```

## Recompilar después de cambios

Si haces cambios en el código fuente (.go), necesitas recompilar:

```bash
# En macOS
./build-windows.sh

# Transfiere el nuevo kahoot-assistant.exe a Windows
```

## Solución de Problemas

### "Permission denied" al ejecutar build-windows.sh

```bash
chmod +x build-windows.sh
```

### Error de compilación

Verifica que tengas Go instalado:
```bash
go version
```

Limpia las dependencias:
```bash
go clean -modcache
go mod download
./build-windows.sh
```

### El .exe no funciona en Windows

1. Verifica que sea Windows 64 bits
2. Asegúrate de que Tesseract esté instalado en Windows
3. Verifica que `config.yaml` esté en el mismo directorio
4. Ejecuta como Administrador si hay problemas de permisos

## Notas Importantes

- ✅ La compilación cruzada es completamente funcional
- ✅ No necesitas Windows para compilar
- ✅ El ejecutable generado funciona en cualquier Windows 64 bits
- ⚠️ Tesseract debe estar instalado **en Windows**, no en macOS
- ⚠️ El programa solo funcionará en Windows, no en macOS

## Compilar para macOS (opcional)

Si también quieres usar el programa en macOS:

```bash
# Para Mac Intel
GOOS=darwin GOARCH=amd64 go build -o kahoot-assistant-mac-intel

# Para Mac Apple Silicon (M1/M2/M3/M4)
GOOS=darwin GOARCH=arm64 go build -o kahoot-assistant-mac-arm
```

**Nota:** Necesitarías instalar Tesseract en macOS:
```bash
brew install tesseract
```

## Estructura de Archivos Recomendada en Windows

```
C:\kahoot-assistant\
├── kahoot-assistant.exe    ← Ejecutable
├── config.yaml              ← Configuración
└── kahoot_log.txt          ← Se crea automáticamente
```

## Siguiente Paso

Lee el archivo `README.md` principal para instrucciones completas de uso.
