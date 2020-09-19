### user service


# installation

    run 'go get' command to install the packages


# database and migrations

    you don't have to worry about DBs, We are using gorm sqlite,
    the moment you run 'go run main.go' command it will create a new db and migrations
    you will later see a file named invoice_service.db created.


# run the application

    $ go run main.go
    this will expose the api on port 9001