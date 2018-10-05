#!/bin/bash
if [ "$AWS_SM_ENV_OPTIONS" != "" ];then
  ENV_OUTPUT=$(aws-sm-env $AWS_SM_ENV_OPTIONS)
  ENV_RC=$?
  if [ "$ENV_RC" -eq "0" ];then
    eval $ENV_OUTPUT
  else
    echo $ENV_OUTPUT
    exit 1
  fi
fi
if [ "$AWS_SM_FILE" != "" -a "$AWS_SM_FILE_OPTIONS" != "" ];then
  FILE_OUTPUT=$(aws-sm-env $AWS_SM_FILE_OPTIONS)
  FILE_RC=$?
  if [ "$FILE_RC" -eq "0" ];then
    echo $FILE_OUTPUT > $AWS_SM_FILE
  else
    echo $FILE_OUTPUT
    exit 1
  fi
fi
exec "$@"
