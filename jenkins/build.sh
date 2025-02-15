#!/bin/bash

# Exportar varibles para compilar aplicación
export $GIN_MODE
export $PORT
export $TYPE
export $PROJECT_ID
export $PRIVATE_KEY_ID
export $PRIVATE_KEY
export $CLIENT_EMAIL
export $CLIENT_ID
export $AUTH_URI
export $TOKEN_URI
export $AUTH_PROVIDER_CERT_URL
export $CLIENT_CERT_URL
export $UNIVERSE_DOMAIN
export $BUCKET

# Manejo de errores: Verifica si go version y go build tienen éxito.
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
