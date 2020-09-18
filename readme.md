### go-microservices


# setting up protoc
you need protoc for protocol buffer
download the release (protoc)
    depending to machine you are using, from https://github.com/protocolbuffers/protobuf/releases
    I have downloaded protoc-3.13.0-linux-x86_32.zip file as i am having linux system
    unzip the file and put the content of the bin folder to your go path go/bin folder so that you can access it like other services,
    you can also set the path its up to you

install grpc curl, ref: https://github.com/fullstorydev/grpcurl
go get github.com/fullstorydev/grpcurl/...

# services
there are two services user and invoice, you can find them in the root directory,
they each have their readme.md file for respective instructions.

# the gateway
use gateway to interact with these services
