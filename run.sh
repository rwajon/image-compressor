#!/bin/bash

mode=${@:-"dev"}

pid=$(lsof -t -i:3000)

[[ $pid ]] && kill -9 $pid
[[ $mode ]] && echo "running $mode server..."

if [[ "$mode" == "dev" ]]; then
  $(go env GOPATH)/bin/air
fi
