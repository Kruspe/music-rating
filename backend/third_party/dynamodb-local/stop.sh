#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_PATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

pidfile="$SCRIPT_PATH/install/pid"
if [ -f "$pidfile" ]; then
  kill -9 "$(cat "$pidfile")" || true
  rm -f "$pidfile"
fi

if ! lsof -ti:8095 &> /dev/null; then
  exit 0
fi

kill -9 "$(lsof -ti:8095)"