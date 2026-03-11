# goDistributedSystem

A minimal distributed task execution prototype in Go with:
- a **master** node exposing:
  - a gRPC service for workers (`:50051`)
  - an HTTP API for submitting tasks (`:8080`)
- a **worker** node that connects to the master, receives queued tasks over a gRPC stream, and simulates task execution.

## Architecture

```text
+----------------------+            gRPC (:50051)             +----------------------+
|       Worker         | <----------------------------------- |       Master         |
|  internal/worker/*   |  stream AssignTask(Request)->Resp    |  internal/master/*   |
|                      |                                      |                      |
|  - ReportStatus() ---+------------------------------------> |  NodeService         |
|  - AssignTask() recv |                                      |  CmdChannel queue    |
+----------------------+                                      +----------+-----------+
                                                                         ^
                                                                         |
                                                              HTTP POST /tasks (:8080)
                                                                         |
                                                            +------------+------------+
                                                            |     API Client (curl)   |
                                                            +-------------------------+
```

The current flow is:
1. Start the master.
2. Start one or more workers.
3. Submit a task to `POST /tasks` on the master.
4. Master enqueues the command and streams it to a connected worker.
5. Worker logs receipt/completion.

Core components:
- `cmd/master/main.go` - boots gRPC + HTTP servers.
- `internal/master/server.go` - gRPC service implementation.
- `internal/master/api.go` - HTTP `/tasks` handler.
- `cmd/worker/main.go` - worker entrypoint.
- `internal/worker/client.go` - worker gRPC client loop.
- `proto/node.proto` - protobuf service/message definitions.
- `pkg/pb/` - generated protobuf/gRPC Go code.

## Prerequisites

- Go `1.24+` (project uses `go 1.24.0` in `go.mod`)
- `protoc` (Protocol Buffers compiler) - only needed if regenerating gRPC code
- Go protobuf plugins (only for regeneration):
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`

Install codegen plugins if needed:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Ensure `$GOPATH/bin` (or `$HOME/go/bin`) is on your `PATH`.

## Setup

```bash
go mod download
```

If you change `proto/node.proto`, regenerate stubs:

```bash
make proto
```

## Run

Start master (terminal 1):

```bash
make run-master
```

Start worker (terminal 2):

```bash
make run-worker
```

## Submit a task

In a third terminal:

```bash
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"command":"echo hello"}'
```

Expected API response:

```json
{"command":"echo hello","status":"queued"}
```

You should then see logs on:
- master: task queued and sent to worker
- worker: task received and completed

## API

### `POST /tasks`

Request body:

```json
{"command":"your-task"}
```

Responses:
- `202 Accepted` on success
- `400 Bad Request` if JSON is invalid or `command` is missing
- `405 Method Not Allowed` for non-POST methods

## Project structure

```text
cmd/
  master/     # master executable
  worker/     # worker executable
internal/
  master/     # master gRPC server + HTTP handlers
  worker/     # worker gRPC client/task loop
pkg/
  pb/         # generated protobuf code
proto/
  node.proto  # gRPC contract
```

## Development notes

- This is a lightweight prototype intended for local development and experimentation.
- Worker processing is currently simulated (sleep + log) in `internal/worker/client.go`.
- `internal/master/queue.go` and `internal/worker/executor.go` are placeholders for future queue/execution abstractions.

## Next improvements

- Add durable task queueing and acknowledgements/retries.
- Add configurable addresses/ports via flags or environment variables.
- Add unit/integration tests.
- Add graceful shutdown and worker heartbeats.
