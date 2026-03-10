# goDistributedSystem

## Protobuf / gRPC code generation

If you need Go protobuf generators, install the two plugins from their correct modules:

```zsh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

`protoc-gen-go` lives in `google.golang.org/protobuf`, but `protoc-gen-go-grpc` does **not**. The gRPC plugin is published from `google.golang.org/grpc/cmd/protoc-gen-go-grpc`.

If your IDE is trying to fetch `google.golang.org/protobuf/cmd/protoc-gen-go-grpc`, update that setting or file watcher to use the `google.golang.org/grpc/...` path instead.
