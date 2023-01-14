# go-grpc-example

# Instructions

1. Install protobuf , grpc 
   
   ```
   brew install protobuf grpc
   ```

2. Install the protocol compiler plugins for Go using the following commands:
  
   ```
   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
   ```

3. Update path 
  ```
  export PATH="$PATH:$(go env GOPATH)/bin"
  ```

4. run 
  ```
   protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative   chat/chat.proto
  ```

5. In one terminal `go run server.go`

6. In another terminal `go run client.go` 




