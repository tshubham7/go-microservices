### user service

# installation
run go get command to install the packages

# database and migrations
make sure you have postgres database created on your local system with the
configuration as defined in package db, user/db/db.go
for eg. i have hard coded the config as
user=postgres, host=localhost, port=5432, password=1234, dbname=user_service_db
migrations will run when you actually run the main file

# run the application
go run main.go
this will expose the api on port 9001