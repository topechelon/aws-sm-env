#!/bin/sh
eval $(aws-sm-env $AWS_SM_ENV_OPTIONS)
exec "$@"
