# Kahoot Assistant

Programa en Go para Windows 64 bits que captura pantalla, extrae texto con OCR, consulta a Gemini AI y envía respuestas por email y/o consola.

## Características

- Captura de pantalla al presionar la tecla Print Screen
- Extracción de texto (OCR) usando Tesseract
- Consulta a Gemini API para responder preguntas de Kahoot
- Respuestas mostradas en consola/CMD
- Envío opcional de pregunta y respuesta por Gmail
- Registro de todas las operaciones en archivo de log
- Procesamiento en menos de 15 segundos
- Sistema de configuración flexible

## Requisitos Previos

### 1. Go 1.21 o superior

Descarga e instala Go desde: https://golang.org/dl/

### 2. Tesseract OCR

**Opción 1: Instalador Windows**
1. Descarga el instalador desde: https://github.com/UB-Mannheim/tesseract/wiki
2. Ejecuta el instalador (recomendado: `tesseract-ocr-w64-setup-5.3.x.exe`)
3. Durante la instalación, asegúrate de incluir los datos de idioma inglés
4. Nota la ruta de instalación (por defecto: `C:\Program Files\Tesseract-OCR`)

**Opción 2: Chocolatey**
```bash
choco install tesseract
```

### 3. Token de Gemini API

1. Ve a [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Inicia sesión con tu cuenta de Google
3. Haz clic en "Create API Key"
4. Copia el token generado

### 4. App Password de Gmail (si usarás email)

1. Ve a tu [Cuenta de Google](https://myaccount.google.com/)
2. Selecciona "Seguridad" en el menú lateral
3. Habilita "Verificación en 2 pasos" (si no está habilitada)
4. Busca "Contraseñas de aplicaciones"
5. Selecciona "Correo" y "Equipo Windows"
6. Genera la contraseña y guárdala (tiene 16 caracteres)

## Instalación

### Paso 1: Clonar o descargar el proyecto

```bash
cd C:\
mkdir kahoot-assistant
cd kahoot-assistant
# Copia todos los archivos del proyecto aquí
```

### Paso 2: Configurar Tesseract

Agrega Tesseract al PATH del sistema:

1. Busca "variables de entorno" en el menú de Windows
2. Haz clic en "Variables de entorno"
3. En "Variables del sistema", selecciona "Path" y haz clic en "Editar"
4. Haz clic en "Nuevo" y agrega: `C:\Program Files\Tesseract-OCR`
5. Haz clic en "Aceptar" en todas las ventanas

### Paso 3: Configurar el archivo config.yaml

Edita el archivo `config.yaml` con tus datos:

```yaml
# Token de Gemini API
gemini_api_key: "TU_TOKEN_DE_GEMINI_AQUI"

# Configuración de correo Gmail
email:
  enabled: true  # Cambiar a false si no quieres usar email
  smtp_host: "smtp.gmail.com"
  smtp_port: 587
  from: "tu_email@gmail.com"
  password: "tu_app_password_de_16_caracteres"
  to: "email_destino@gmail.com"

# Configuración de salida
output:
  console: true  # Mostrar en consola
  email: true    # Enviar por email
  log_file: true # Guardar en archivo de log

# Resto de la configuración (puedes dejar los valores por defecto)
```

### Paso 4: Descargar dependencias

```bash
go mod tidy
go mod download
```

### Paso 5: Compilar el programa

```bash
# Para Windows 64 bits
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o kahoot-assistant.exe .
```

O en una sola línea:
```bash
go build -ldflags="-H windowsgui" -o kahoot-assistant.exe .
```

## Uso

### Iniciar el programa

```bash
.\kahoot-assistant.exe
```

O simplemente haz doble clic en `kahoot-assistant.exe`

### Funcionamiento

1. El programa se quedará esperando a que presiones la tecla **Print Screen**
2. Cuando presiones Print Screen:
   - Captura la pantalla completa
   - Extrae el texto visible usando OCR
   - Envía la pregunta a Gemini AI
   - Muestra la respuesta en la consola
   - Opcionalmente envía la pregunta y respuesta por email
   - Guarda todo en el archivo de log
3. Todo el proceso toma menos de 15 segundos
4. El programa seguirá ejecutándose y esperando más capturas
5. Presiona **Ctrl+C** para salir

### Ejemplo de salida

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

Esperando que presiones la tecla Print Screen...

Tecla Print Screen detectada!

--- Iniciando procesamiento ---
1. Capturando pantalla...
   Captura guardada en: C:\Users\...\temp\kahoot_screenshot_xxx.png
2. Extrayendo texto con OCR...
   Texto extraído exitosamente

   PREGUNTA DETECTADA:
   What is the capital of France?

3. Consultando a Gemini AI...
   Respuesta recibida de Gemini

╔════════════════════════════════════════════════════════════════╗
║                        RESPUESTA                               ║
╚════════════════════════════════════════════════════════════════╝
The capital of France is Paris.
═══════════════════════════════════════════════════════════════

4. Enviando correo...
   ✓ Correo enviado a: destino@gmail.com
5. Guardando en log...
   ✓ Guardado en: kahoot_log.txt

✓ Procesamiento completado en 8.42 segundos

Listo para la siguiente captura.
Presiona Print Screen nuevamente...
```

## Configuración Avanzada

### Desactivar el envío de email

En `config.yaml`:
```yaml
email:
  enabled: false
```

O:
```yaml
output:
  email: false
```

### Cambiar el idioma de OCR

Si tus preguntas están en otro idioma, cambia en `config.yaml`:
```yaml
ocr:
  language: "spa"  # Para español
  # Otros: fra (francés), deu (alemán), etc.
```

Debes instalar el paquete de idioma correspondiente en Tesseract.

### Ajustar timeouts

En `config.yaml`:
```yaml
timeouts:
  max_processing_time: 20  # Aumentar si necesitas más tiempo
  gemini_timeout: 15
  email_timeout: 5
```

## Solución de Problemas

### Error: "Tesseract not found"

- Verifica que Tesseract esté instalado
- Asegúrate de que esté en el PATH del sistema
- Reinicia la terminal/CMD después de agregar al PATH
- O especifica la ruta completa en config.yaml:
  ```yaml
  ocr:
    tesseract_path: "C:\\Program Files\\Tesseract-OCR\\tesseract.exe"
  ```

### Error: "Invalid API key"

- Verifica que hayas copiado correctamente el token de Gemini
- Asegúrate de no tener espacios adicionales
- Genera un nuevo token si es necesario

### Error al enviar email

- Verifica que uses una App Password, no tu contraseña normal de Gmail
- Asegúrate de que la verificación en 2 pasos esté habilitada
- Verifica tu conexión a internet
- Intenta con otro email de destino

### El OCR no extrae texto correctamente

- Asegúrate de que el texto en pantalla sea claro y legible
- Aumenta la resolución de tu pantalla si es posible
- Verifica que estés capturando la región correcta

### Error: "Access denied" al presionar Print Screen

- Ejecuta el programa como Administrador
- Click derecho en kahoot-assistant.exe > "Ejecutar como administrador"

## Estructura del Proyecto

```
kahoot-assistant/
├── config.yaml           # Configuración
├── config.go            # Carga de configuración
├── screenshot.go        # Captura de pantalla
├── ocr.go               # Extracción de texto OCR
├── gemini.go            # Cliente Gemini API
├── email.go             # Envío de correos
├── logger.go            # Sistema de logging
├── keyboard_windows.go  # Detección de teclas Windows
├── main.go              # Programa principal
├── go.mod               # Dependencias
└── README.md            # Este archivo
```

## Notas de Seguridad

- **NUNCA** compartas tu archivo `config.yaml` con tus tokens y contraseñas
- Agrega `config.yaml` a `.gitignore` si usas Git
- Los tokens de API y contraseñas son sensibles
- Usa variables de entorno en producción

## Licencia

Este proyecto es de código abierto para uso educativo.

## Soporte

Si encuentras problemas:
1. Verifica la sección de Solución de Problemas
2. Revisa que todos los requisitos estén instalados
3. Verifica los logs en `kahoot_log.txt`
