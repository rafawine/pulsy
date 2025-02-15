#!/bin/bash

# Variables importantes al principio y en mayúsculas para mejor visibilidad
SERVICES_DIR="$HOME/workspace/services"
SERVICE_FILE="$SERVICES_DIR/$JOB_NAME.service"
LN_SERVICES_DIR="/etc/systemd/system"
LN_SERVICE_FILE="$SERVICES_LN_DIR/$JOB_NAME.service"

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
  if file_already_exists "$LN_SERVICE_FILE"; then
    return 0  # Código de éxito
  fi

  if ! command -v ln &> /dev/null; then
    echo "Error: El comando 'ln' no está instalado." >&2
    return 1
  fi

  sudo ln -s "$SERVICE_FILE" "$LN_SERVICES_DIR"

  if [ $? -eq 0 ]; then  # Verifica el código de salida de ln
    echo "Enlace simbólico de servicio $JOB_NAME.service creado exitosamente."
  else
    echo "Error: Falló la creación del enlace simbólico."
    return 1  # Código de error
  fi

  sudo systemctl daemon-reload

  if [ $? -eq 0 ]; then  # Verifica el código de salida de systemctl
    echo "Systemctl recargado exitosamente."
  else
    echo "Error: Systemctl no recargado."
    return 1  # Código de error
  fi

  return 0  # Código de éxito
}

# Verificación y creación del directorio (con manejo de errores)
if [ ! -d "$SERVICES_DIR" ]; then
  if ! mkdir -p "$SERVICES_DIR"; then  # -p crea directorios padres si es necesario
    echo "Error: No se pudo crear el directorio $SERVICES_DIR."
    exit 1 # Salida del script con error
  fi
  echo "Directorio $SERVICES_DIR creado exitosamente."
fi

if create_service; then
  echo "Servicio $JOB_NAME creado exitosamente."
else
  echo "Error: Servicio $JOB_NAME ya existente."
fi

if create_ln_service; then
  echo "Enlace simbólico de servicio $JOB_NAME creado exitosamente."
else
  echo "Error: Enlace simbólico de servicio $JOB_NAME ya existente."
fi

exit 0 # Salida del script con éxito
