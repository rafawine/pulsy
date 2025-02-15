#!/bin/bash

# Nombre del servicio
service_name="$JOB_NAME"

# Verifica si el servicio está activo
service_status=$(systemctl is-active "$service_name")

# Evalúa el estado del servicio
if [[ "$service_status" == "active" ]]; then
  echo "El servicio $service_name está activo. Reiniciando..."
  sudo systemctl restart "$service_name"
elif [[ "$service_status" == "inactive" ]]; then
  echo "El servicio $service_name está inactivo. Iniciando..."
  sudo systemctl start "$service_name"
else
  echo "No se pudo determinar el estado del servicio $service_name."
fi