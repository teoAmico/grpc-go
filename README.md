# gRpc Go Boilerplate

## Setup the project

Execute the following cmd to resolve the versions:

`go mod tidy`

`go mod vendor`

https://github.com/grpc-ecosystem/grpc-gateway

## Health Servce Check

https://github.com/grpc/grpc/blob/master/doc/health-checking.md

grpcurl -plaintext 127.0.0.1:9090 grpc.health.v1.Health/Check

## Remove listener on port

sudo lsof -i :9090

kill -9 'PID'
