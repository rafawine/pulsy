#!/bin/bash

STATUS=$(systemctl is-active $JOB_NAME)

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

sudo systemctl status "$JOB_NAME"; then
  echo "Error: No se pudo obtener el estatus del servicio $JOB_NAME."
  return 1
fi

exit 0
