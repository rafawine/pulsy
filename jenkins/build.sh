#!/bin/bash

# Manejo de errores: Verifica si go version y go build tienen éxito.
if ! go version &> /dev/null; then
  echo "Error: go no está instalado o no está en el PATH."
  exit 1
fi

if [ -f "$HOME/workspace/envs/${JOB_NAME}_env.sh" ]; then
  source "$HOME/workspace/envs/${JOB_NAME}_env.sh"

  if [ ! $? -eq 0 ]; then  # Verifica el código de salida de source
    echo "Error: Falló al cargar las variables de entorno."
    return 1  # Código de error
  fi
else
  echo "Error: Variables de entorno no encontradas."
  exit 1 
fi

if ! go build -o ./tmp/$JOB_NAME ./cmd/server; then
  echo "Error: Falló la compilación de la aplicación."
  exit 1
fi

echo "Aplicación compilada exitosamente"

exit 0 # Salida del script con éxito
