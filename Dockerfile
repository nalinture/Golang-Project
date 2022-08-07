FROM golang:alpine
COPY . /bookingapp
WORKDIR /bookingapp
CMD go run learn.go
