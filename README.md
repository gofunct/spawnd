# Spawnd

## Roadmap
### Transport Generation
- [ ] gRPC server
- [ ] http Server
- [ ] http-grpc proxy server

### Debug Generation
- [ ] prometheus metrics and handler
- [ ] jaeger tracing
- [ ] pprof handler
- [ ] readiness handler
- [ ] liveness handler
- [ ] prometheus

### Project Generation
- [ ] makefile
- [ ] cobra cli
- [ ] config file
- [ ] .gitignore

### Deployment
- [ ] prometheus dockerfile
- [ ] jaeger dockerfile
- [ ] endpoints dockerfiles

## Usage

```text
Usage:
  spawnd [command]

Available Commands:
  cobra       generate a cobra client command
  convert     A brief description of your command
  help        Help about any command
  init        Initialize a GrpcGen Application
  protoc      n/a
  proxy       A brief description of your command
  server      Create a new command to start a grpc and http based application

Flags:
  -h, --help             help for spawnd
      --service string   The protobuf message used for this configuration

```