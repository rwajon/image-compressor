#!/bin/bash

mode=${@:-"dev"}
if [[ -f .env ]]; then
  port=$(grep -w .env -e 'PORT' | sed 's/PORT=//')
  pid=$(sudo lsof -t -i :$port)
fi

if [[ $pid ]]; then
  echo -e "Port $port address already in use \nEnter a new port: "
  read answer
fi
echo $answer

if ! [[ -f $(go env GOPATH)/bin/air ]]; then
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
fi

if [[ "$mode" == "dev" ]]; then
  PORT=$answer $(go env GOPATH)/bin/air
fi
