# grpc-golang
Studying gRPC protocol with Golang

> **Note:**
> To install gRPC in your PC and implement other functionalities in this (or other) code, you need to must be installed the library of protoc generator. To install, follow the instructions on topic [Install protoc generator][installation]

### To run :rocket:
```sh
  # the server gRPC
  $ go run cmd/server/index.go

  # the client gRPC
  $ go run cmd/client/index.go
```

### To install proto-buffer-generator
| Linux | ```sudo apt install protobuf-compiler``` or ```snap install protobuf --classic``` |
|-------|-----------------------------------------------------------------------------------|
| MAC   | ```brew install protobuf # Mac, required Homebrew installed```                    |

and to finish... run:
```sh 
  $ go get \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

<!-- Links -->
[installation]: #to-install-proto-buffer-generator
