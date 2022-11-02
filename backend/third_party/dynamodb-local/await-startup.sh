#!/usr/bin/env bash

counter=0

until [ "$dynamodbStatus" == "com.amazonaws.dynamodb.v20120810#MissingAuthenticationToken" ]
do
  dynamodbStatus=$(curl -s "http://localhost:8095" | jq -r '.__type')
  echo "Waiting for local dynamodb to start"
  ((counter++))
  sleep 1
  if [ "$counter" -gt 30 ]; then
      echo "local dynamodb did not start"
      exit 1
  fi
done

echo "local dynamodb running"