#!/bin/bash

# Variables importantes al principio y en mayúsculas para mejor visibilidad
SERVICES_DIR="$HOME/workspace/services"
SERVICE_FILE="$SERVICES_DIR/$JOB_NAME.service"
SERVICE_LN_FILE="/etc/systemd/system/$JOB_NAME.service"

# Verifica si un archivo existe.
#
# Args:
#   path_file: La ruta al archivo.
#
# Returns:
#   0 si el archivo no existe.
#   1 si el archivo existe o no se proporcionó una ruta.
file_already_exists() {
  local path_file="$1"

  # Verifica si se proporcionó una ruta.
  if [ -z "$path_file" ]; then
    echo "Error: Se debe proporcionar la ruta del archivo." >&2
    return 1
  fi

  # Verifica si el archivo ya existe.
  if [ -f "$path_file" ]; then
    echo "Error: El archivo '$path_file' ya existe." >&2
    return 0
  fi

  return 1  # Código si el archivo no existe.
}

# Función principal (mejor modularización)
create_service() {
  # Verifica si el archivo ya existe (salida temprana)
  if file_already_exists "$SERVICE_FILE"; then
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

create_ln_service() {
  # Verifica si el archivo ya existe (salida temprana)
  if file_already_exists "$SERVICE_LN_FILE"; then
    return 0  # Código de éxito
  fi

  echo "Se crea enlace simbolico"
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