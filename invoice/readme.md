### invoice service


### installing grpc and protoc
ref: http://grpc.github.io/docs/quickstart/go.html
go get -u google.golang.org/grpc

download the release (protoc)
    depending to machine you are using, from https://github.com/protocolbuffers/protobuf/releases
    I have downloaded protoc-3.13.0-linux-x86_32.zip file as i am having linux system
    unzip the file and put the content of the bin folder to your go path go/bin folder so that you can access it like other services,
    you can also set the path its up to you

install grpc curl, ref: https://github.com/fullstorydev/grpcurl
go get github.com/fullstorydev/grpcurl/...


# installation
run go get command to install the packages

# database and migrations
make sure you have postgres database created on your local system with the
configuration as defined in package db, invoice/db/db.go
for eg. i have hard coded the config as
user=postgres, host=localhost, port=5432, password=1234, dbname=invoice_service_db
migrations will run when you actually run the main file

# run the application
go run main.go
this will expose the grpc server on port 9002

# test grpc
to test grpc you can use makefile, check out that file