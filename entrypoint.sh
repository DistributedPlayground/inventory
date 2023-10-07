#!/bin/sh
export $(grep -v '^#' .env | xargs)

make generate

# run app
if [ "$DEBUG_MODE" = "true" ]; then
  echo "----- DEBUG_MODE is true"
  air -c .air-debug.toml
else
  echo "----- DEBUG_MODE is false"
  air
fi