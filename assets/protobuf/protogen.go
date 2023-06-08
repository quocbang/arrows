package protobuf

// !TCP generate region.
//go:generate protoc -I=. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../pkg/protobuf api/tcp/tcp.proto
//go:generate protoc -I=. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true,allow_delete_body=true:../../pkg/protobuf api/tcp/tcp.proto
//go:generate protoc -I=. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --swagger_out=allow_delete_body=true,logtostderr=true:../../pkg/protobuf api/tcp/tcp.proto

// !UDP henerate region.
//go:generate protoc -I=. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../pkg/protobuf api/udp/udp.proto
//go:generate protoc -I=. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true,allow_delete_body=true:../../pkg/protobuf api/udp/udp.proto
//go:generate protoc -I=. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --swagger_out=allow_delete_body=true,logtostderr=true:../../pkg/protobuf api/udp/udp.proto

// TODO: Request.
//go:generate protoc -I=. --go_out=paths=source_relative:../../pkg/protobuf authentication/user.proto
