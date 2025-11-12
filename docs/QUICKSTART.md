# Guía Rápida - Kahoot Assistant

## Para Usuarios de Windows

### 1. Instalar Tesseract OCR

1. Descarga: https://github.com/UB-Mannheim/tesseract/wiki
2. Ejecuta: `tesseract-ocr-w64-setup-5.3.x.exe`
3. Durante instalación: Marca "English language data"
4. Agrega al PATH del sistema: `C:\Program Files\Tesseract-OCR`

**Verificar instalación:**
```cmd
tesseract --version
```

### 2. Obtener Token de Gemini

1. Ve a: https://makersuite.google.com/app/apikey
2. Inicia sesión con Google
3. Click en "Create API Key"
4. Copia el token (algo como: `AIzaSyC...`)

### 3. Configurar Email (Opcional)

Si quieres recibir respuestas por email:

1. Ve a: https://myaccount.google.com/security
2. Activa "Verificación en 2 pasos"
3. Busca "Contraseñas de aplicaciones"
4. Genera una contraseña (16 caracteres)
5. Guárdala

### 4. Configurar el Programa

Edita `config.yaml`:

```yaml
gemini_api_key: "PEGA_AQUI_TU_TOKEN_DE_GEMINI"

email:
  enabled: true  # Cambiar a false si no quieres email
  from: "tu_email@gmail.com"
  password: "contraseña_de_16_caracteres"
  to: "donde_quieres_recibir@gmail.com"
```

### 5. Ejecutar

Doble click en `kahoot-assistant.exe`

O desde CMD:
```cmd
kahoot-assistant.exe
```

### 6. Usar

1. Abre Kahoot en tu navegador
2. Ejecuta `kahoot-assistant.exe`
3. Cuando aparezca una pregunta, presiona **Print Screen**
4. Espera 5-15 segundos
5. Ve la respuesta en la consola (y en tu email si lo configuraste)

## Para Usuarios de macOS (Compilar)

### 1. Compilar para Windows

```bash
cd /ruta/a/kahoot-assistant
chmod +x build-windows.sh
./build-windows.sh
```

### 2. Transferir a Windows

- Copia `kahoot-assistant.exe` a tu PC Windows
- Copia `config.example.yaml` y renómbralo a `config.yaml`
- Sigue los pasos de "Para Usuarios de Windows"

## Configuración Mínima

Si solo quieres ver respuestas en consola (sin email):

```yaml
gemini_api_key: "TU_TOKEN"

email:
  enabled: false  # Desactivar email

output:
  console: true
  email: false
  log_file: true
```

## Ejemplo de Uso

```
===========================================
   Kahoot Assistant - Iniciando...
===========================================

Configuración cargada correctamente
- Salida en consola: true
- Envío de email: false
- Guardar en log: true

Programa listo. Presiona Print Screen.

[Presionas Print Screen]

--- Iniciando procesamiento ---
1. Capturando pantalla...
2. Extrayendo texto con OCR...
   PREGUNTA DETECTADA:
   What is the capital of France?

3. Consultando a Gemini AI...

╔════════════════════════════════════════╗
║            RESPUESTA                   ║
╚════════════════════════════════════════╝
The capital of France is Paris.
═══════════════════════════════════════

✓ Procesamiento completado en 6.8 segundos

Listo para la siguiente captura.
```

## Solución Rápida de Problemas

| Error | Solución |
|-------|----------|
| "Tesseract not found" | Instala Tesseract y agrégalo al PATH |
| "Invalid API key" | Verifica tu token de Gemini |
| "Error enviando correo" | Usa App Password, no contraseña normal |
| No extrae texto | Verifica que Tesseract esté instalado |
| Programa no responde | Ejecuta como Administrador |

## Recursos

- Token Gemini: https://makersuite.google.com/app/apikey
- Tesseract Windows: https://github.com/UB-Mannheim/tesseract/wiki
- App Password Gmail: https://myaccount.google.com/apppasswords
- Documentación completa: Ver `README.md`
- Instalación detallada: Ver `INSTALL.md`
- Compilar desde macOS: Ver `README_macOS.md`

## Tips

1. **Velocidad**: El programa tarda 5-15 segundos. Ten paciencia.
2. **Claridad**: Asegúrate que el texto en pantalla sea legible
3. **Pantalla completa**: Mejor capturar en pantalla completa
4. **Log**: Revisa `kahoot_log.txt` si hay problemas
5. **Seguridad**: Nunca compartas tu `config.yaml`

## Atajos de Teclado

- `Print Screen`: Capturar y analizar
- `Ctrl + C`: Salir del programa

---

**¡Listo!** Ahora puedes usar Kahoot Assistant para obtener respuestas rápidas a preguntas de Kahoot.
