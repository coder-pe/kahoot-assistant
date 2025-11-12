# Guía de Instalación Rápida - Kahoot Assistant

## Paso 1: Instalar requisitos

### 1.1 Instalar Go
1. Descarga desde: https://golang.org/dl/
2. Instala la última versión (1.21 o superior)
3. Verifica en CMD: `go version`

### 1.2 Instalar Tesseract OCR
1. Descarga desde: https://github.com/UB-Mannheim/tesseract/wiki
2. Ejecuta el instalador `tesseract-ocr-w64-setup-5.3.x.exe`
3. Durante la instalación, marca "Additional language data (download)" > English
4. Anota la ruta de instalación (normalmente `C:\Program Files\Tesseract-OCR`)

### 1.3 Agregar Tesseract al PATH
1. Presiona `Win + R`, escribe `sysdm.cpl` y presiona Enter
2. Ve a la pestaña "Opciones avanzadas"
3. Haz clic en "Variables de entorno"
4. En "Variables del sistema", selecciona "Path" y haz clic en "Editar"
5. Haz clic en "Nuevo" y agrega: `C:\Program Files\Tesseract-OCR`
6. Haz clic en "Aceptar" en todas las ventanas
7. **Cierra y vuelve a abrir tu terminal/CMD**

## Paso 2: Obtener tokens y credenciales

### 2.1 Obtener token de Gemini API
1. Ve a: https://makersuite.google.com/app/apikey
2. Inicia sesión con tu cuenta de Google
3. Haz clic en "Create API Key"
4. Copia y guarda el token

### 2.2 Obtener App Password de Gmail (si usarás email)
1. Ve a: https://myaccount.google.com/
2. Haz clic en "Seguridad" en el menú lateral
3. Activa "Verificación en 2 pasos" (si no está activa)
4. Busca "Contraseñas de aplicaciones" (App Passwords)
5. Selecciona "Correo" y "Computadora Windows"
6. Genera y copia la contraseña de 16 caracteres

## Paso 3: Configurar el proyecto

### 3.1 Abrir CMD en el directorio del proyecto
```bash
cd C:\ruta\a\kahoot-assistant
```

### 3.2 Crear archivo de configuración
```bash
copy config.example.yaml config.yaml
```

### 3.3 Editar config.yaml
Abre `config.yaml` con un editor de texto y completa:

```yaml
gemini_api_key: "PEGA_TU_TOKEN_DE_GEMINI_AQUI"

email:
  enabled: true  # Cambiar a false si no quieres usar email
  from: "tu_email@gmail.com"
  password: "tu_app_password_de_16_caracteres"
  to: "email_destino@gmail.com"
```

**Guarda el archivo.**

## Paso 4: Compilar

En CMD, ejecuta:
```bash
build.bat
```

O si prefieres PowerShell:
```powershell
.\build.ps1
```

Deberías ver el mensaje "Compilación exitosa!"

## Paso 5: Ejecutar

```bash
.\kahoot-assistant.exe
```

O simplemente:
```bash
run.bat
```

## Verificación

Si todo está correcto, deberías ver:

```
===========================================
   Kahoot Assistant - Iniciando...
===========================================

Configuración cargada correctamente
- Salida en consola: true
- Envío de email: true
- Guardar en log: true
- Tiempo máximo de procesamiento: 15 segundos

Programa listo. Presiona Print Screen para capturar y analizar.
Presiona Ctrl+C para salir.
```

## Problemas Comunes

### "Tesseract not found"
- Verifica que Tesseract esté en el PATH
- Reinicia tu CMD/terminal
- O especifica la ruta completa en config.yaml:
  ```yaml
  ocr:
    tesseract_path: "C:\\Program Files\\Tesseract-OCR\\tesseract.exe"
  ```

### "Invalid API key"
- Verifica que copiaste correctamente el token de Gemini
- No debe tener espacios adicionales
- Debe estar entre comillas en config.yaml

### Error de email
- Verifica que uses App Password, no tu contraseña normal
- Asegúrate de que la verificación en 2 pasos esté activa
- Revisa que el email sea correcto

### "go: command not found"
- Go no está instalado o no está en el PATH
- Reinicia tu terminal después de instalar Go
- Verifica con: `go version`

## Listo!

Ahora puedes usar el programa. Presiona **Print Screen** cuando tengas una pregunta de Kahoot en pantalla y recibirás la respuesta en segundos.
