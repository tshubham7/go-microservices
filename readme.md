### go-microservices

# key points!

    Golang
    Protocol Buffers
    GRPC
    Gorm
    Gin
    Restful API
    Swager
    Docker -> find the docker section below to use it


# direct way to test?

    please find the docker section below.


# What's in it?

    We have two services user and invoice, with their separate database 
    their endpoints are exposed separately, 
    
    We have a gateway so that you don't have to keep track using different endpoints,
    you will be using only one host, which is gateway's endpoint.

    There is an inter service communication through GRPC between user and invoice services,
    each time there is an alteration happen in the database table record(create, delete, update),
    the invoice is created using the invoice service.

    We are also keeping track of the invoice service logs in database. (event sourcing)


# protocol buffer, protoc

    protoc is a protocol buffer compiler(https://developers.google.com/protocol-buffers/docs/gotutorial), 

    if you want to have protoc compiler for protocol buffer
    download the release (protoc)
    depending to machine you are using, from https://github.com/protocolbuffers/protobuf/releases
    I have downloaded protoc-3.13.0-linux-x86_32.zip file as i am having linux system.
    unzip the file and put the content of the bin folder to your go path go/bin folder so that you can access it like other services,
    you can also set the path its up to you


# services

    find the following services
    1. User service
    2. Invoice service
    3. The Gateway to interact with these services


# how to start?

    run the gateway server
    run the invoice server
    run the user server
    

# docker images

    download the docker images from the following urls
    gateway:
        $ docker pull tshubham7/go-microservices-gateway:latest
    run using command
        $ docker run -it -p 8080:8080 tshubham7/go-microservices-gateway

    invoice service:
        $ docker pull tshubham7/go-microservices-invoice-service:latest
    run using command:
        $ docker run -it -p 9002:9002 -p 5000:5000 tshubham7/go-microservices-invoice-service

    user service:
        $ docker pull tshubham7/go-microservices-user-service:latest
    run using command:
        $ docker run -it -p 9001:9001 tshubham7/go-microservices-user-service

    
    note: 
        if you are getting 502 error while using these images, 
        use http://172.17.0.1:8080/ as host if you are running docker images 
        for eg. http://172.17.0.1:8080/api/invoice?sort=created_at&order=desc&limit=10&offset=0


# quick test

    run this on your browser http://127.0.0.1:8080/api/invoice?sort=created_at&order=desc&limit=10&offset=0
    use http://172.17.0.1:8080/ if you are getting 502


# endpoints/api doc

    find the swagger documentation for restful api at dir. specs/swagger/api.yaml
    copy the content of this file, then open https://editor.swagger.io/# and paste the content. 

    find the postman collection for restful api at dir. specs/postman/collection.json
    download the file, open the your postman tool then import this json file into it.

