#!/bin/bash

mode=${@:-"dev"}

[[ $mode ]] && echo "running $mode server..."

if [[ "$mode" == "dev" ]]; then
  $(go env GOPATH)/bin/air
fi
