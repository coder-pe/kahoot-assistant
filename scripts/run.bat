@echo off
echo ============================================
echo      Kahoot Assistant - Iniciando...
echo ============================================
echo.

REM Verificar que existe el archivo de configuracion
if not exist "config.yaml" (
    echo Error: No se encuentra el archivo config.yaml
    echo.
    echo Por favor, crea tu archivo de configuracion:
    echo   1. Copia config.example.yaml a config.yaml
    echo   2. Edita config.yaml con tus datos
    echo.
    pause
    exit /b 1
)

REM Verificar que existe el ejecutable
if not exist "kahoot-assistant.exe" (
    echo Error: No se encuentra kahoot-assistant.exe
    echo.
    echo Por favor, compila el programa primero:
    echo   .\build.bat
    echo.
    pause
    exit /b 1
)

REM Ejecutar el programa
kahoot-assistant.exe
