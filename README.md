## for generating the go code files from the .proto

protoc --go_out=. --go-grpc_out=. proto/greet.proto

## to remove all the errors from the gen files

go mod tidy

## start the server and client on two seperate terminals :

go run \*.go
