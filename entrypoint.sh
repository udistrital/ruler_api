#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export RULER_DB_USER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/ruler_api/db/username --output text --query Parameter.Value)"
  export RULER_DB_PASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/ruler_api/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
