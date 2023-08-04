#!/bin/bash

tsServices=("profile")
goServices=("authentication" "account" )
authInternalServices=("account")
gatewayServices=("authentication" "account")

function cleanup() {
  folderPath="$1/generated"
  echo "$folderPath"
  if [ -d "$folderPath" ]; then
    rm -r "$folderPath"
  fi
  mkdir -p "$folderPath"
}

function generateTS() {
  if [ -d "$1" ]; then
    cd "$1" || return
    npm run protoc
    cd ..
  fi
}

function generateGo() {
  path="$1/generated"
  protoc --go_out="$path" --go-grpc_out="$path" --proto_path=proto "$2.proto"
}

# generate code for TS services
for service in "${tsServices[@]}"; do
  generateTS "$service" "$service"
  echo "generated proto for TS service: $service"
done

# generate code for Go services
for service in "${goServices[@]}"; do
  cleanup "$service"
  if [ "$service" == "authentication" ]; then
    # generate code for authentication
    for innerService in "${authInternalServices[@]}"; do
      generateGo "$service" "$innerService"
      echo "generated proto for authentication: $innerService"
    done
  fi
  generateGo "$service" "$service"
  echo "generated proto for Go service: $service"
done

# generate code for gateway
cleanup api-gateway
for service in "${gatewayServices[@]}"; do
  generateGo api-gateway "$service"
  echo "generated proto for api-gateway: $service"
done

sleep .5
