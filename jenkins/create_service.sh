#!/bin/bash

if [ -d /var/lib/jenkins/workspace/services ]; then
  
else
  mkdir /var/lib/jenkins/workspace/services
  echo "services folder created"
fi

create_service() {
  if [ -f /var/lib/jenkins/workspace/services/$JOB_NAME ]; then
    echo "nothing to do"
  else
    descripcion="file transfer api"
    working_directory="path/path"
    user="jenkins"
    execute="path/path"

    sed -e "s/@DESCRIPCION@/$descripcion/" -e "s/@WORK_DIR@/$working_directory/" -e "s/@USR@/$user/" -e "s/@EXEC@/$execute/" ./jenkins/systemd.template > /var/lib/jenkins/workspace/services/$JOB_NAME.service
  fi
}
