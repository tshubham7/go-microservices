FROM golang:1.13

LABEL maintainer="Shubham Dhanera<tshubham19@gmail.com>"
LABEL tag="go-microservices-invoice-service"

WORKDIR /src
# Copy over the app files
COPY . /src

RUN go build -o main

EXPOSE 9002 5000

CMD ["./main"]
