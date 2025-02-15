#!/bin/bash

# Verifica si el servicio está activo
STATUS=$(systemctl is-active $JOB_NAME)

# Evalúa el estado del servicio
if [[ "$STATUS" == "active" ]]; then
  if ! sudo systemctl restart "$JOB_NAME" > /dev/null 2>&1; then
    echo "Error: No se pudo reiniciar el servicio $JOB_NAME."
    return 1
  fi
elif [[ "$STATUS" == "inactive" ]]; then
  if ! sudo systemctl start "$JOB_NAME" > /dev/null 2>&1; then
    echo "Error: No se pudo reiniciar el servicio $JOB_NAME."
    return 1
  fi
else
  echo "Error: No se pudo determinar el estado del servicio $JOB_NAME."
fi

exit 0