The project *grpc_stringsvc* implements GRPC with a GOLANG project.

#Prerequisites
go get golang.org/x/net/context
go get google.golang.org/grpc

As well as this packages to build your RPC library to share between server and client
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto

#Project
This repository contains a string server and a client to consume it and it uses the GRPC protocol based on protobuf 3. Few methods are implemented as UpperCase and LowerCase. 
aaaaaaa
#Generates .pb
Everytime you wish to modify your contract, run the following command. If the change is breaking, you client and server will have to be updated
`protoc -I stringsvc_contract/ stringsvc_contract/stringsvc_contract.proto --go_out=plugins=grpc:stringsvc_contract`

{{ff
