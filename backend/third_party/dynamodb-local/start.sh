#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_PATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

"$SCRIPT_PATH/stop.sh"

cd "$SCRIPT_PATH/install"
java \
  -Djava.library.path=./DynamoDBLocal_lib \
  -jar DynamoDBLocal.jar \
  -inMemory \
  -sharedDb \
  -port 8095 >/dev/null &

echo "$!" > "$SCRIPT_PATH/install/pid"