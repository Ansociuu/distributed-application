Scientific RPC calculator (Go)

Files created in d:\\distributed_application:
- server.go  (RPC server)
- client.go  (RPC client example)

Run:
1) In terminal A: cd d:\\distributed_application && go run server.go
2) In terminal B: cd d:\\distributed_application && go run client.go

Or build:
cd d:\\distributed_application
go build -o server.exe server.go
go build -o client.exe client.go
then run server.exe (terminal A) and client.exe (terminal B).

The client demonstrates Add, Div, Pow and Sqrt calls to the RPC server on 127.0.0.1:1234.
