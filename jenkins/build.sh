#!/bin/bash

if ! go version &> /dev/null; then
  echo "Error: go no está instalado o no está en el PATH."
  exit 1
fi

if ! go build -o ./tmp/$JOB_NAME ./cmd/server; then
  echo "Error: Falló la compilación."
  exit 1
fi

echo "Compilación realizada correctamente"

exit 0
