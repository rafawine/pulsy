#!/bin/bash

# Variables importantes al principio y en mayúsculas para mejor visibilidad
SERVICES_DIR="/var/lib/jenkins/workspace/services"
TEMPLATE_FILE="./jenkins/systemd.template"  # Ruta relativa al script
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
  working_directory="path/path"
  user="jenkins"
  execute="path/path"

  # Verifica que el archivo template exista
    if [ ! -f "$TEMPLATE_FILE" ]; then
        echo "Error: Archivo de plantilla $TEMPLATE_FILE no encontrado."
        return 1 # Código de error
    fi

  # Uso de sed para crear el archivo .service (con manejo de errores)
  if ! sed -e "s/@DESCRIPCION@/$descripcion/" -e "s/@WORK_DIR@/$working_directory/" -e "s/@USR@/$user/" -e "s/@EXEC@/$execute/" "$TEMPLATE_FILE" > "$SERVICE_FILE"; then
    echo "Error: Falló la creación del archivo $SERVICE_FILE."
    return 1 # Código de error
  fi

  echo "Servicio $JOB_NAME creado exitosamente en $SERVICE_FILE."
  return 0 # Código de éxito
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