#!/bin/bash

# Variables importantes al principio y en mayúsculas para mejor visibilidad
SERVICES_DIR="$HOME/workspace/services"
SERVICE_FILE="$SERVICES_DIR/$JOB_NAME.service"

# Función principal (mejor modularización)
create_service() {
  # Verifica si el archivo ya existe (salida temprana)
  if [ -f "$SERVICE_FILE" ]; then
    echo "Servicio $JOB_NAME ya existe. No se hará nada."
    return 0  # Código de éxito
  fi

  # Definición de variables (podrían ser argumentos de la función)
  descripcion="file transfer api"
  working_directory="$WORKSPACE"
  user="$USER"
  execute="$WORKSPACE/tmp/$JOB_NAME"

  cat << EOF > "$SERVICE_FILE"
[Unit]
Description="$descripcion"
After=network.target

[Service]
Type=simple
WorkingDirectory="$working_directory"
User="$user"
ExecStart="$execute"
Restart=always

[Install]
WantedBy=multi-user.target
EOF

  if [ $? -eq 0 ]; then  # Verifica el código de salida de cat
    echo "Archivo de servicio $JOB_NAME.service creado exitosamente."
    return 0  # Código de éxito
  else
    echo "Error: Falló la creación del archivo de servicio."
    return 1  # Código de error
  fi
}

# Verificación y creación del directorio (con manejo de errores)
if [ ! -d "$SERVICES_DIR" ]; then
  if ! mkdir -p "$SERVICES_DIR"; then  # -p crea directorios padres si es necesario
    echo "Error: No se pudo crear el directorio $SERVICES_DIR."
    exit 1 # Salida del script con error
  fi
  echo "Directorio $SERVICES_DIR creado."
fi

# Llamada a la función principal y manejo del código de retorno
if create_service; then
  echo "Proceso completado."
else
  echo "El proceso falló."
  exit 1 # Salida del script con error
fi

exit 0 # Salida del script con éxito