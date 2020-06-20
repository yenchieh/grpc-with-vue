# gRPC + Go + Vue + Typescript
* It's a voting website that combines gRPC, Go, Vue, and Typescript.
-  just a personal practice project

To run the project, you will need to run `Go - Backend`, `Vue -frontend`, and `Envoy`.

## gRPC - Proto File
* Located in `proto` folder
* [vote.proto](proto/vote.proto)
To generate Go
```
# Generate files under server/gen/pb
$ make gen-go-proto
```
To generate Typescript
```
# Generate files under web/gen/pb
$ make gen-ts-proto
```

## Backend Go
* [grpc](https://github.com/grpc/grpc-go)
* [protobuf](https://github.com/golang/protobuf)

To install libs for backend Go
```
# Under server folder
$ make deps
```
Run the server (it's on 5110 port)
```
$ make dev
  go run main.go
  2020/06/20 17:41:59 Server on 5110

```

## Frontend Vue
* [Vue](https://github.com/vuejs/vue)
* [grpc-web](https://github.com/grpc/grpc-web)

To install libs
```
# Under web folder
$ yarn
```
To run frontend web server
```
$ yarn dev
```
## Envoy
It needs Envoy between our client and server for transparent translation between HTTP/1.x and HTTP/2.

To build envoy image
```
# under envoy folder
$ make build-image
```

To run envoy in docker
```
$ make run-image
```