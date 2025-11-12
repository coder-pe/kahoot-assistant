# Scripts de Compilación - Kahoot Assistant

Este directorio contiene scripts para compilar el proyecto para diferentes plataformas.

## 📋 Scripts Disponibles

### 🔨 Compilación Modular (Recomendado)

#### `build-current.sh` - Compilar para tu plataforma actual
Detecta automáticamente tu sistema y compila solo para él.

```bash
./scripts/build-current.sh
```

**Detecta automáticamente:**
- macOS Intel → `kahoot-assistant-darwin-amd64`
- macOS Apple Silicon → `kahoot-assistant-darwin-arm64`
- Linux x64 → `kahoot-assistant-linux-amd64`
- Linux ARM64 → `kahoot-assistant-linux-arm64`

---

#### `build-windows.sh` - Solo Windows
Compila únicamente para Windows 64 bits.

```bash
./scripts/build-windows.sh
```

**Genera:** `builds/kahoot-assistant-windows-amd64.exe`

---

#### `build-macos.sh` - Solo macOS
Compila para macOS (puedes elegir arquitectura).

```bash
# Ambas arquitecturas (Intel y Apple Silicon)
./scripts/build-macos.sh

# Solo Intel
./scripts/build-macos.sh intel

# Solo Apple Silicon (M1/M2/M3/M4)
./scripts/build-macos.sh arm
```

**Genera:**
- `builds/kahoot-assistant-darwin-amd64` (Intel)
- `builds/kahoot-assistant-darwin-arm64` (Apple Silicon)

---

#### `build-linux.sh` - Solo Linux
Compila para Linux (puedes elegir arquitectura).

```bash
# Ambas arquitecturas (x64 y ARM64)
./scripts/build-linux.sh

# Solo x64
./scripts/build-linux.sh amd64

# Solo ARM64 (Raspberry Pi, etc.)
./scripts/build-linux.sh arm64
```

**Genera:**
- `builds/kahoot-assistant-linux-amd64` (x64)
- `builds/kahoot-assistant-linux-arm64` (ARM64)

---

### 🚀 Compilación Completa

#### `build-all.sh` - Todas las plataformas
Compila para TODAS las plataformas soportadas.

```bash
./scripts/build-all.sh
```

**Genera 5 binarios:**
1. `kahoot-assistant-windows-amd64.exe` (Windows 64)
2. `kahoot-assistant-darwin-amd64` (macOS Intel)
3. `kahoot-assistant-darwin-arm64` (macOS M1/M2/M3/M4)
4. `kahoot-assistant-linux-amd64` (Linux x64)
5. `kahoot-assistant-linux-arm64` (Linux ARM64)

---

## 📁 Ubicación de Binarios

Todos los binarios compilados se guardan en:
```
builds/
├── kahoot-assistant-windows-amd64.exe
├── kahoot-assistant-darwin-amd64
├── kahoot-assistant-darwin-arm64
├── kahoot-assistant-linux-amd64
└── kahoot-assistant-linux-arm64
```

---

## 🎯 Casos de Uso

### Desarrollo en macOS M4
```bash
# Compilar solo para tu Mac
./scripts/build-current.sh

# O específicamente Apple Silicon
./scripts/build-macos.sh arm

# Ejecutar
./builds/kahoot-assistant-darwin-arm64
```

### Desarrollo en Linux
```bash
# Compilar solo para Linux
./scripts/build-current.sh

# O compilar para distribuir en ambas arquitecturas
./scripts/build-linux.sh
```

### Compilación cruzada para Windows (desde macOS/Linux)
```bash
# Solo Windows
./scripts/build-windows.sh

# Transferir builds/kahoot-assistant-windows-amd64.exe a Windows
```

### Preparar release para todas las plataformas
```bash
# Compilar todo
./scripts/build-all.sh

# Los 5 binarios estarán listos en builds/
```

---

## ⚡ Comparación

| Script | Plataformas | Tiempo | Uso |
|--------|-------------|--------|-----|
| `build-current.sh` | 1 (actual) | ~10s | Desarrollo rápido |
| `build-windows.sh` | 1 (Windows) | ~10s | Solo Windows |
| `build-macos.sh` | 1-2 (macOS) | ~10-20s | Solo macOS |
| `build-linux.sh` | 1-2 (Linux) | ~10-20s | Solo Linux |
| `build-all.sh` | 5 (todas) | ~50s | Release completo |

---

## 🛠️ Opciones Avanzadas

### Compilación manual con Go
Si prefieres compilar manualmente:

```bash
# Para tu plataforma actual
go build -o kahoot-assistant ./cmd/kahoot-assistant

# Para Windows desde macOS/Linux
GOOS=windows GOARCH=amd64 go build -o builds/kahoot-assistant.exe ./cmd/kahoot-assistant

# Para macOS Apple Silicon desde Linux/Windows
GOOS=darwin GOARCH=arm64 go build -o builds/kahoot-assistant-mac ./cmd/kahoot-assistant

# Con optimizaciones (reduce tamaño)
go build -ldflags="-s -w" -o kahoot-assistant ./cmd/kahoot-assistant
```

---

## 📝 Notas

- **Permisos**: Los scripts necesitan permisos de ejecución (`chmod +x`)
- **Go requerido**: Todos los scripts requieren Go 1.21+
- **Directorio builds/**: Se crea automáticamente si no existe
- **Limpieza**: Los scripts limpian builds anteriores automáticamente
- **Tiempo**: Cada compilación toma ~10 segundos

---

## 🔄 Flujo de Trabajo Recomendado

### Durante Desarrollo
```bash
# Rápido: solo tu plataforma
./scripts/build-current.sh
```

### Antes de Commit
```bash
# Verificar que compila en todas las plataformas
./scripts/build-all.sh
```

### Para Distribuir
```bash
# Compilar todo y comprimir
./scripts/build-all.sh
cd builds
tar -czf kahoot-assistant-v1.0.0.tar.gz kahoot-assistant-*
```

---

## 🐛 Solución de Problemas

### "Permission denied"
```bash
chmod +x scripts/*.sh
```

### "go: command not found"
Instala Go 1.21+ desde https://golang.org/dl/

### Compilación falla
```bash
# Limpiar y reintentar
go clean -cache
go mod tidy
./scripts/build-current.sh
```

---

## 📖 Más Información

- **Documentación principal**: `docs/INDEX.md`
- **README del proyecto**: `README.md` (raíz)
- **Guía multi-plataforma**: `docs/README_MULTIPLATFORM.md`
