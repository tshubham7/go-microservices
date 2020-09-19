FROM golang:1.13

LABEL maintainer="Shubham Dhanera<tshubham19@gmail.com>"
LABEL tag="go-microservices-user-service"

WORKDIR /src
# Copy over the app files
COPY . /src

RUN go build -o main

EXPOSE 9001

CMD ["./main"]
