GRPC

File grpc go_out

    protoc --go_out=plugins=grpc:. ./proto/*.proto

File grpc geteway_out

    protoc --grpc-gateway_out=proto ./proto/*.proto