#!/bin/bash

printf "\033[92m" # print texts in green

while [ $# -gt 0 ]; do
  case "$1" in
  --os=*)
    os="${1#*=}"
    ;;
  --arch=*)
    arch="${1#*=}"
    ;;
  --app_name=*)
    app_name="${1#*=}"
    ;;
  *)
    printf "***************************\n"
    printf "* Error(get_args.sh): Invalid argument (${1}).\n"
    printf "***************************\n"
    exit 1
    ;;
  esac
  shift
done

os=${os:-"linux"}
arch=${arch:-"amd64"}
app_name="image_compressor"

echo "env GOOS=$os GOARCH=$arch go build -o ${app_name}_${arch} ."
env GOOS=$os GOARCH=$arch go build -o ${app_name}_${arch} .
