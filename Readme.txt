Install GRPC:
And Get all sample:
go get google.golang.org/grpc


go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto

Generates .pb
protoc -I stringsvc_contract/ stringsvc_contract/stringsvc_contract.proto --go_out=plugins=grpc:stringsvc_contract