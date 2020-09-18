### invoice service


# installing grpc and protoc

    find the instructions about this topic here, https://github.com/tshubham7/go-microservices/blob/master/readme.md     



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