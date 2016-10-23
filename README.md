The project *grpc_stringsvc* implements a GRPC service in a GOLANG project.

#Prerequisites

##Install GRPC
`go get google.golang.org/grpc`

##Install Protocol Buffers v3 
Download pre-compiled binaries https://github.com/google/protobuf/releases and update your PATH to include the path to the protoc binary file.<br />
<br />
Next install the protoc plugin for Go to build your RPC library that you will then share between server and client<br />
`go get -u github.com/golang/protobuf/protoc-gen-go`<br />
`go get -u github.com/golang/protobuf/proto`<br />
<br />
About prerequisite if something is missing check the official documentation http://www.grpc.io/docs/quickstart/go.html <br />
 
#Project 
This repository contains a string server, a client and a GRPC service. Communication is based on proto buffer v3. The GRPC service defines few methods as UpperCase and LowerCase.  
 
#Generates .pb 
Everytime you wish to modify your service, run the following command. Then compile your server and client to ensure they still are compliant with our GRPC service.<br />
`protoc -I stringsvc_contract/ stringsvc_contract/stringsvc_contract.proto --go_out=plugins=grpc:stringsvc_contract` 
