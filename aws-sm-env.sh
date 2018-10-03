#!/bin/sh
OUTPUT=$(aws-sm-env $AWS_SM_ENV_OPTIONS)
if [ "$?" -eq "0" ];then
  eval $OUTPUT
  exec "$@"
else
  echo $OUTPUT
fi
