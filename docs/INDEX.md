# Kahoot Assistant - Índice de Documentación

## 🚀 Inicio Rápido

**¿Primera vez? Empieza aquí:**

1. **[QUICKSTART.md](QUICKSTART.md)** - Guía rápida de 5 minutos
2. **[README_MULTIPLATFORM.md](README_MULTIPLATFORM.md)** - Guía completa multi-plataforma ⭐

## 📖 Por Tipo de Usuario

### Soy usuario de Windows
1. Leer: **[README.md](README.md)** o **[INSTALL.md](INSTALL.md)**
2. Descargar: `kahoot-assistant-windows-amd64.exe`
3. Instalar: Tesseract OCR
4. Configurar: `config.yaml`
5. Ejecutar y usar

### Soy usuario de macOS
1. Leer: **[README_macOS.md](README_macOS.md)** o **[README_MULTIPLATFORM.md](README_MULTIPLATFORM.md)**
2. Descargar: `kahoot-assistant-darwin-arm64` (M1/M2/M3/M4) o `kahoot-assistant-darwin-amd64` (Intel)
3. Instalar: `brew install tesseract`
4. Configurar: `config.yaml`
5. Ejecutar y usar

### Soy usuario de Linux
1. Leer: **[README_MULTIPLATFORM.md](README_MULTIPLATFORM.md)**
2. Descargar: `kahoot-assistant-linux-amd64`
3. Instalar: `sudo apt install tesseract-ocr` (o tu gestor de paquetes)
4. Configurar: `config.yaml`
5. Ejecutar y usar

### Quiero compilar desde código
1. Leer: **[README_macOS.md](README_macOS.md)** (si estás en Mac)
2. Ejecutar: `./build-all.sh` para todas las plataformas
3. O: `go build` para tu plataforma actual

## 📚 Documentación Completa

### Guías de Usuario

| Archivo | Descripción | Para quién |
|---------|-------------|-----------|
| **[QUICKSTART.md](QUICKSTART.md)** | Inicio rápido | Todos - Primera lectura |
| **[README_MULTIPLATFORM.md](README_MULTIPLATFORM.md)** | Guía multi-plataforma completa | Todos - Lectura principal |
| **[README.md](README.md)** | Documentación principal | Windows principalmente |
| **[README_ES.md](README_ES.md)** | README en español | Hispanohablantes |
| **[INSTALL.md](INSTALL.md)** | Instalación paso a paso | Principiantes |
| **[README_macOS.md](README_macOS.md)** | Compilar en macOS | Usuarios Mac |

### Documentación Técnica

| Archivo | Descripción |
|---------|-------------|
| **[PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)** | Resumen técnico del proyecto |
| **[INDEX.md](INDEX.md)** | Este archivo - Índice |

### Archivos de Configuración

| Archivo | Descripción |
|---------|-------------|
| `config.yaml` | Tu configuración personal (NO compartir) |
| `config.example.yaml` | Plantilla de configuración |

## 🛠️ Scripts

### Scripts de Compilación

| Archivo | Plataforma | Descripción |
|---------|-----------|-------------|
| `build-all.sh` | macOS/Linux | Compila para TODAS las plataformas ⭐ |
| `build-windows.sh` | macOS/Linux | Solo Windows 64 bits |
| `build.bat` | Windows | Compilar en Windows (CMD) |
| `build.ps1` | Windows | Compilar en Windows (PowerShell) |
| `run.bat` | Windows | Ejecutar en Windows |

### Uso de Scripts

```bash
# Compilar para todas las plataformas (macOS/Linux)
./build-all.sh

# Compilar solo para Windows desde macOS
./build-windows.sh

# En Windows (CMD)
build.bat

# En Windows (PowerShell)
.\build.ps1
```

## 📦 Binarios

### Disponibles después de compilar

| Archivo | Plataforma | Tamaño |
|---------|-----------|--------|
| `kahoot-assistant-windows-amd64.exe` | Windows 64 | 15 MB |
| `kahoot-assistant-darwin-amd64` | macOS Intel | 14 MB |
| `kahoot-assistant-darwin-arm64` | macOS M1/M2/M3/M4 | 14 MB |
| `kahoot-assistant-linux-amd64` | Linux 64 | 14 MB |
| `kahoot-assistant-linux-arm64` | Linux ARM64 | 13 MB |

## 📂 Estructura del Proyecto

```
kahoot-assistant/
├── 📘 Documentación
│   ├── INDEX.md (este archivo)
│   ├── README_MULTIPLATFORM.md ⭐
│   ├── QUICKSTART.md
│   ├── README.md
│   ├── README_ES.md
│   ├── README_macOS.md
│   ├── INSTALL.md
│   └── PROJECT_SUMMARY.md
│
├── 💻 Código Fuente
│   ├── main.go (principal)
│   ├── config.go
│   ├── gemini.go
│   ├── email.go
│   ├── logger.go
│   ├── ocr.go
│   ├── process_common.go
│   ├── screenshot_common.go
│   │
│   ├── Windows específico
│   │   ├── keyboard_windows.go
│   │   ├── screenshot_windows.go
│   │   └── process_windows.go
│   │
│   ├── macOS específico
│   │   ├── keyboard_darwin.go
│   │   ├── screenshot_darwin.go
│   │   └── process_darwin.go
│   │
│   └── Linux específico
│       ├── keyboard_linux.go
│       ├── screenshot_linux.go
│       └── process_linux.go
│
├── 🛠️ Scripts
│   ├── build-all.sh ⭐
│   ├── build-windows.sh
│   ├── build.bat
│   ├── build.ps1
│   └── run.bat
│
├── ⚙️ Configuración
│   ├── config.yaml (tu config)
│   ├── config.example.yaml (plantilla)
│   ├── go.mod
│   └── go.sum
│
└── 📦 Binarios (después de compilar)
    ├── kahoot-assistant-windows-amd64.exe
    ├── kahoot-assistant-darwin-amd64
    ├── kahoot-assistant-darwin-arm64
    ├── kahoot-assistant-linux-amd64
    └── kahoot-assistant-linux-arm64
```

## 🎯 Flujo de Lectura Recomendado

### Para comenzar rápido
```
1. QUICKSTART.md
2. README_MULTIPLATFORM.md (sección de tu OS)
3. ¡Configurar y usar!
```

### Para instalación detallada
```
1. INSTALL.md
2. README_MULTIPLATFORM.md
3. Configurar config.yaml
4. Probar el programa
```

### Para desarrolladores
```
1. PROJECT_SUMMARY.md
2. README_macOS.md (si vas a compilar)
3. Revisar código fuente
4. ./build-all.sh
```

## 🔗 Enlaces Importantes

### Recursos Externos
- **Gemini API Token:** https://makersuite.google.com/app/apikey
- **Tesseract Windows:** https://github.com/UB-Mannheim/tesseract/wiki
- **App Password Gmail:** https://myaccount.google.com/apppasswords
- **Homebrew (macOS):** https://brew.sh/

### Comandos Rápidos

**Instalar Tesseract:**
```bash
# macOS
brew install tesseract

# Ubuntu/Debian
sudo apt install tesseract-ocr

# Fedora
sudo dnf install tesseract

# Arch
sudo pacman -S tesseract
```

**Compilar:**
```bash
# Todas las plataformas
./build-all.sh

# Solo tu plataforma
go build
```

**Ejecutar:**
```bash
# Windows
.\kahoot-assistant-windows-amd64.exe

# macOS (M1/M2/M3/M4)
./kahoot-assistant-darwin-arm64

# macOS (Intel)
./kahoot-assistant-darwin-amd64

# Linux
./kahoot-assistant-linux-amd64
```

## ❓ Preguntas Frecuentes

### ¿Qué archivo debo leer primero?
👉 **[QUICKSTART.md](QUICKSTART.md)** o **[README_MULTIPLATFORM.md](README_MULTIPLATFORM.md)**

### ¿Cómo instalo en Windows?
👉 **[INSTALL.md](INSTALL.md)** o **[README.md](README.md)**

### ¿Cómo compilo en macOS?
👉 **[README_macOS.md](README_macOS.md)**

### ¿Funciona en Linux?
👉 Sí! Lee **[README_MULTIPLATFORM.md](README_MULTIPLATFORM.md)**

### ¿Cómo obtengo el token de Gemini?
👉 https://makersuite.google.com/app/apikey

### ¿Necesito saber programar?
👉 No, solo descarga el binario para tu sistema y configura `config.yaml`

## 📊 Comparación de Archivos de Documentación

| Archivo | Longitud | Nivel | Idioma |
|---------|----------|-------|--------|
| QUICKSTART.md | Corta | Principiante | Inglés/Español |
| README_MULTIPLATFORM.md | Larga | Intermedio | Español |
| README.md | Larga | Intermedio | Inglés |
| README_ES.md | Media | Todos | Español |
| INSTALL.md | Media | Principiante | Español |
| README_macOS.md | Media | Avanzado | Español |
| PROJECT_SUMMARY.md | Larga | Avanzado | Español |

## ✅ Checklist de Inicio

- [ ] Leí QUICKSTART.md o README_MULTIPLATFORM.md
- [ ] Descargué el binario para mi sistema
- [ ] Instalé Tesseract OCR
- [ ] Obtuve mi token de Gemini API
- [ ] Configuré config.yaml con mis datos
- [ ] (Opcional) Configuré App Password de Gmail
- [ ] Probé el programa con una pregunta

## 🎉 ¿Listo para Empezar?

👉 **[QUICKSTART.md](QUICKSTART.md)** - ¡Comienza aquí!

👉 **[README_MULTIPLATFORM.md](README_MULTIPLATFORM.md)** - Guía completa

---

**Proyecto:** Kahoot Assistant Multi-Plataforma
**Versión:** 1.0.0
**Plataformas:** Windows, macOS, Linux
**Estado:** ✅ Listo para producción
