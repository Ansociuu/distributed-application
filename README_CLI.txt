Interactive CLI client for Calculator RPC

Files:
- cli_client.go (interactive client)

Run:
1) Start server in terminal A:
   cd d:\\distributed_application
   go run server.go

2) In terminal B run CLI client:
   cd d:\\distributed_application
   go run cli_client.go

Optional: specify server address:
   go run cli_client.go -addr 127.0.0.1:1234

Examples inside CLI:
> add 2 3
> pow 2 8
> sqrt 16
> sin 1.5708

Notes:
- Trigonometric functions expect radians.
- Build exe: go build -o cli_client.exe cli_client.go
