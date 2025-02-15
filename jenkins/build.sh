#!/bin/bash

# Test variables de entorno:

echo $GIN_MODE
echo $PORT
echo $POPROJECT_IDRT

# 1. Manejo de errores: Verifica si go version y go build tienen éxito.
if ! go version &> /dev/null; then
  echo "Error: go no está instalado o no está en el PATH."
  exit 1
fi

if ! go build -o ./tmp/$JOB_NAME ./cmd/server; then
  echo "Error: Falló la compilación de la aplicación."
  exit 1
fi

echo "Aplicación compilada exitosamente"

exit 0 # Salida del script con éxito
